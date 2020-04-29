package core

import (
	"github.com/cpacia/openbazaar3.0/database"
	"github.com/cpacia/openbazaar3.0/events"
	"github.com/cpacia/openbazaar3.0/models"
	peer "github.com/libp2p/go-libp2p-core/peer"
	"testing"
	"time"
)

func TestFollowerTracker_ConnectDisconnect(t *testing.T) {
	node, err := MockNode()
	if err != nil {
		t.Fatal(err)
	}

	defer node.DestroyNode()

	startSub, err := node.SubscribeEvent(&events.TrackerStarted{})
	if err != nil {
		t.Fatal(err)
	}

	ft := NewFollowerTracker(node.repo, node.eventBus, node.ipfsNode.PeerHost)
	ft.Start()

	select {
	case <-startSub.Out():
	case <-time.After(time.Second * 10):
		t.Fatal("Timeout waiting on channel")
	}

	connectSub, err := node.SubscribeEvent(&events.TrackerPeerConnected{})
	if err != nil {
		t.Fatal(err)
	}

	p, err := peer.Decode("QmfQkD8pBSBCBxWEwFSu4XaDVSWK6bjnNuaWZjMyQbyDub")
	if err != nil {
		t.Fatal(err)
	}
	node.eventBus.Emit(&events.Follow{PeerID: p.Pretty()})
	node.eventBus.Emit(&events.PeerConnected{Peer: p})

	select {
	case <-connectSub.Out():
	case <-time.After(time.Second * 10):
		t.Fatal("Timeout waiting on channel")
	}

	if _, ok := ft.connected[p]; !ok {
		t.Error("Peer is disconnected")
	}

	disconnectSub, err := node.SubscribeEvent(&events.TrackerPeerDisconnected{})
	if err != nil {
		t.Fatal(err)
	}

	node.eventBus.Emit(&events.PeerDisconnected{Peer: p})

	select {
	case <-disconnectSub.Out():
	case <-time.After(time.Second * 10):
		t.Fatal("Timeout waiting on channel")
	}

	if _, ok := ft.connected[p]; ok {
		t.Error("Peer is connected")
	}

	var stat models.FollowerStat
	err = node.repo.DB().View(func(tx database.Tx) error {
		return tx.Read().First(&stat).Error
	})
	if err != nil {
		t.Fatal(err)
	}

	if stat.PeerID != p.Pretty() {
		t.Errorf("Incorrect peer ID. Expected %s, got %s", stat.PeerID, p.Pretty())
	}

	if stat.ConnectedDuration == 0 {
		t.Error("Failed to record connection duration")
	}
}

func TestFollowerTracker_ConnectToFollowers(t *testing.T) {
	mocknet, err := NewMocknet(5)
	if err != nil {
		t.Fatal(err)
	}

	defer mocknet.TearDown()

	var followers models.Followers
	for _, node := range mocknet.Nodes()[1:] {
		followers = append(followers, node.Identity().Pretty())
	}

	err = mocknet.Nodes()[0].repo.DB().Update(func(tx database.Tx) error {
		return tx.SetFollowers(followers)
	})
	if err != nil {
		t.Fatal(err)
	}

	ft := NewFollowerTracker(mocknet.Nodes()[0].repo, mocknet.Nodes()[0].eventBus, mocknet.Nodes()[0].ipfsNode.PeerHost)
	ft.Start()

	ft.tryConnectFollowers()

	if len(ft.connected) != 4 {
		t.Errorf("Incorrect number of connected followers. Expected %d, got %d", 4, len(ft.connected))
	}
}
