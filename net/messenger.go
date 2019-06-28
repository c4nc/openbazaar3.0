package net

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"errors"
	"github.com/cpacia/openbazaar3.0/models"
	"github.com/cpacia/openbazaar3.0/net/pb"
	"github.com/cpacia/openbazaar3.0/repo"
	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/ptypes"
	"github.com/jinzhu/gorm"
	peer "github.com/libp2p/go-libp2p-peer"
	"sync"
	"time"
)

// RetryInterval is the interval at which retry sending messages
// that haven't yet been ACKed.
const (
	RetryInterval = time.Minute * 10
	SendTimeout   = time.Minute
)

// Messenger manages the reliable sending of outgoing messages.
// New messages are saved to the database and continually retried
// until the recipient receives it.
type Messenger struct {
	ns        *NetworkService
	repo      *repo.Repo
	ctx       context.Context
	ctxCancel context.CancelFunc
	mtx       sync.RWMutex
}

// NewMessenger returns a Messenger and starts the retry service.
func NewMessenger(ns *NetworkService, repo *repo.Repo) *Messenger {
	ctx, cancel := context.WithCancel(context.Background())
	m := &Messenger{ns, repo, ctx, cancel, sync.RWMutex{}}
	go m.startRetryingMessages()
	return m
}

// Stop shuts down the Messenger.
func (m *Messenger) Stop() {
	m.ctxCancel()
}

// ReliablySendMessage persists the message to the database before sending, then continually retries
// the send until it finally goes through.
func (m *Messenger) ReliablySendMessage(tx *gorm.DB, peer peer.ID, message *pb.Message, done chan<- struct{}) error {
	m.mtx.Lock()
	defer m.mtx.Unlock()

	ser, err := proto.Marshal(message)
	if err != nil {
		return err
	}

	// Before we do anything save the message to the database. This way
	// we can retry sending the message until we know for sure that it
	// has been delivered.
	err = tx.Save(&models.OutgoingMessage{
		ID:                message.MessageID,
		Recipient:         peer.Pretty(),
		SerializedMessage: ser,
		LastAttempt:       time.Now(),
	}).Error
	if err != nil {
		return err
	}

	// Then send the message
	go m.trySendMessage(peer, message, done)

	return nil
}

// ProcessACK deletes the message from the database after it has been
// ACKed so we no longer try sending.
func (m *Messenger) ProcessACK(tx *gorm.DB, message *pb.Message) error {
	if message.MessageType != pb.Message_ACK {
		return errors.New("message is not type ACK")
	}
	ack := new(pb.AckMessage)
	if err := ptypes.UnmarshalAny(message.Payload, ack); err != nil {
		return err
	}
	log.Debugf("Received ACK for message ID %s", ack.AckedMessageID)
	m.mtx.Lock()
	defer m.mtx.Unlock()

	return tx.Delete(&models.OutgoingMessage{ID: ack.AckedMessageID}).Error
}

// SendACK sends an ACK for the message with the given ID to the provided
// peer. The ACK send is only attempted just once and unlike other messages
// is not persisted to the database. It is expect that the message handler
// will send an ACK for every duplicate message it receives. This implies
// that the sender will continue sending messages until he receives an
// ACK and the recipient will continue ACKing them until he stops receiving
// duplicate messages.
func (m *Messenger) SendACK(messageID string, peer peer.ID) {
	log.Debugf("Sending ACK for message ID: %s", messageID)

	ack := &pb.AckMessage{
		AckedMessageID: messageID,
	}

	payload, err := ptypes.MarshalAny(ack)
	if err != nil {
		log.Errorf("Error marshalling ack message: %s", err)
		return
	}

	mid := make([]byte, 20)
	rand.Read(mid)

	msg := &pb.Message{
		MessageID:   base64.StdEncoding.EncodeToString(mid),
		MessageType: pb.Message_ACK,
		Payload: payload,
	}
	go m.trySendMessage(peer, msg, nil)
}

func (m *Messenger) startRetryingMessages() {
	// Run once at startup
	go m.retryAllMessages()

	// Then every RetryInterval
	ticker := time.NewTicker(RetryInterval)
	for range ticker.C {
		select {
		case <-m.ctx.Done():
			ticker.Stop()
			return
		default:
		}
		go m.retryAllMessages()
	}
}

// trySendMessage tries to send the message directly to the peer using a
// network connection. If that fails, it sends the message over the offline
// messaging system.
func (m *Messenger) trySendMessage(peer peer.ID, message *pb.Message, done chan<- struct{}) {
	defer func() {
		if done != nil {
			close(done)
		}
	}()

	ctx, cancel := context.WithTimeout(context.Background(), SendTimeout)
	defer cancel()

	if err := m.ns.SendMessage(ctx, peer, message); err != nil {
		log.Debugf("Failed to connect to peer %s. Sending offline message.", peer)
		// We failed to deliver directly to the peer. Let's send
		// using the offline system.

		// TODO:
		return
	}
	log.Debugf("Message %s direct send successful", message.MessageID)
}

// retryAllMessages loads all un-ACKed messages from the database and
// tries to send them again.
func (m *Messenger) retryAllMessages() {
	m.mtx.RLock()
	var messages []models.OutgoingMessage
	err := m.repo.DBRead(func(tx *gorm.DB) error {
		return tx.Find(&messages).Error
	})
	if err != nil {
		log.Error("Error loading outgoing messages from the database: %s", err)
		m.mtx.RUnlock()
		return
	}
	m.mtx.RUnlock()

	for _, message := range messages {
		pmes := new(pb.Message)
		if err := proto.Unmarshal(message.SerializedMessage, pmes); err != nil {
			log.Error("Error unmarshalling outgoing message: %s", err)
			continue
		}
		pid, err := peer.IDFromString(message.Recipient)
		if err != nil {
			log.Error("Error parsing peer ID in outgoing message: %s", err)
			continue
		}
		go m.trySendMessage(pid, pmes, nil)

		err = m.repo.DBUpdate(func(tx *gorm.DB) error {
			return tx.Model(&message).Update("last_attempt", time.Now()).Error
		})
		if err != nil {
			log.Error("Error updating last attempt for outgoing message: %s", err)
		}

	}
}
