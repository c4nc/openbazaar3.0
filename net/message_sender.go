package net

import (
	"context"
	"fmt"
	"github.com/cpacia/openbazaar3.0/net/pb"
	ggio "github.com/gogo/protobuf/io"
	inet "github.com/libp2p/go-libp2p-core/network"
	"github.com/libp2p/go-libp2p-core/peer"
	"sync"
	"time"
)

type messageSender struct {
	s         inet.Stream
	w         ggio.WriteCloser
	r         ggio.ReadCloser
	lk        sync.Mutex
	p         peer.ID
	ns        *NetworkService
	singleMes int
	invalid   bool
}

var ReadMessageTimeout = time.Minute * 5
var ErrContextDone = fmt.Errorf("write context closed")

func (ns *NetworkService) messageSenderForPeer(ctx context.Context, p peer.ID) (*messageSender, error) {
	ns.msMtx.Lock()
	ms, ok := ns.messageSenders[p]
	if ok {
		// Already have a messageSender for this peer
		// so we can just return it.
		ns.msMtx.Unlock()
		return ms, nil
	}

	// messageSender doesn't exist for this peer so we'll
	// create a new one and attempt to open a new stream
	// with them.
	ms = &messageSender{p: p, ns: ns}
	ns.messageSenders[p] = ms
	ns.msMtx.Unlock()

	if err := ms.ctxPrepOrInvalidate(ctx); err != nil {

		// If we error here it could be because we hit a race
		// condition where another messageSender was opened while
		// we were trying to open this one. If so we'll just return
		// the new one.
		ns.msMtx.Lock()
		defer ns.msMtx.Unlock()

		if msCur, ok := ns.messageSenders[p]; ok {
			// Changed. Use the new one, old one is invalid and
			// not in the map so we can just throw it away.
			if ms != msCur {
				return msCur, nil
			}
			// Not changed, remove the now invalid stream from the
			// map.
			delete(ns.messageSenders, p)
		}
		// Invalid but not in map. Must have been removed by a disconnect.
		return nil, err
	}
	// All ready to go.
	return ms, nil
}

// invalidate is called before this messageSender is removed from the strmap.
// It prevents the messageSender from being reused/reinitialized and then
// forgotten (leaving the stream open).
func (ms *messageSender) invalidate() {
	ms.invalid = true
	if ms.s != nil {
		ms.s.Reset()
		ms.s = nil
	}
}

func (ms *messageSender) ctxPrepOrInvalidate(ctx context.Context) error {
	ms.lk.Lock()
	defer ms.lk.Unlock()

	errCh := make(chan error)
	go func() {
		errCh <- ms.prep()
	}()

	select {
	case err := <-errCh: // Prep finished. Check error.
		if err != nil {
			ms.invalidate()
		}
		return err
	case <-ctx.Done(): // Context finished. Invalidate.
		ms.invalidate()
		return ErrContextDone
	case <-ms.ns.ctx.Done(): // Network Service context finished. Invalidate.
		ms.invalidate()
		return ErrContextDone
	}
}

func (ms *messageSender) prep() error {
	if ms.invalid {
		return fmt.Errorf("message sender has been invalidated")
	}
	if ms.s != nil {
		return nil
	}

	nstr, err := ms.ns.host.NewStream(ms.ns.ctx, ms.p, ms.ns.protocolID)
	if err != nil {
		return err
	}

	ms.r = ggio.NewDelimitedReader(nstr, inet.MessageSizeMax)
	ms.w = ggio.NewDelimitedWriter(nstr)
	ms.s = nstr

	return nil
}

// streamReuseTries is the number of times we will try to reuse a stream to a
// given peer before giving up and reverting to the old one-message-per-stream
// behavior.
const streamReuseTries = 3

func (ms *messageSender) sendMessage(ctx context.Context, pmes *pb.Message) error {
	ms.lk.Lock()
	defer ms.lk.Unlock()
	retry := false
	for {
		if err := ms.prep(); err != nil {
			return err
		}

		if err := ms.ctxWriteMsg(ctx, pmes); err != nil {
			ms.s.Reset()
			ms.s = nil

			if err == ErrContextDone {
				return err
			}

			if retry {
				log.Debug("error writing message, bailing: ", err)
				return err
			}
			log.Debug("error writing message, trying again: ", err)
			retry = true
			continue
		}

		if ms.singleMes > streamReuseTries {
			ms.s.Close()
			ms.s = nil
		} else if retry {
			ms.singleMes++
		}

		return nil
	}
}

func (ms *messageSender) ctxWriteMsg(ctx context.Context, pmes *pb.Message) error {
	errCh := make(chan error)
	go func() {
		errCh <- ms.w.WriteMsg(pmes)
	}()

	select {
	case err := <-errCh:
		return err
	case <-ctx.Done():
		return ErrContextDone
	}
}
