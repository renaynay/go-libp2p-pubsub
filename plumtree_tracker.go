package pubsub

import (
	"fmt"
	"github.com/libp2p/go-libp2p/core/host"
	"github.com/libp2p/go-libp2p/core/peer"
	"sync"
)

const (
	DefaultFanout = 10
)

type Sampler interface {
	FindPeers(count int) []peer.ID
	NeighborDown(peer.ID)
	NeighborUp(peer.ID)
}

type tracker struct {
	host    host.Host
	sampler Sampler

	eagerMu sync.RWMutex
	eager   map[peer.ID]struct{}

	lazyMu sync.RWMutex
	lazy   map[peer.ID]struct{}
}

func newTracker(host host.Host, sampler Sampler) *tracker {
	return &tracker{
		host:    host,
		sampler: sampler,
		eager:   make(map[peer.ID]struct{}),
		lazy:    make(map[peer.ID]struct{}),
	}
}

func (t *tracker) start() {
	log.Debug("starting node", "peerID", t.host.ID())
	eagerPeers := t.sampler.FindPeers(DefaultFanout)
	t.eagerMu.Lock()
	defer t.eagerMu.Unlock()
	for _, p := range eagerPeers {
		log.Debug("initializing with eager peer", "peer", p.String())
		t.eager[p] = struct{}{}
	}
}

// onGRAFT attempts to graft the given peer to the eager list
// for eager pushing, and removes the given peer from the lazy list
// if it exists.
func (t *tracker) onGRAFT(p peer.ID) error {
	return t.moveToEager(p)
}

// onPRUNE attempts to prune the given peer from the eager list
// and adds the given peer to the lazy list if it doesn't
// already exist.
func (t *tracker) onPRUNE(p peer.ID) error {
	return t.moveToLazy(p)
}

// moveToEager moves the given peer to the eager peerset
func (t *tracker) moveToEager(p peer.ID) error {
	log.Debug("moving peer to eager set", "peer", p.String())
	t.lazyMu.Lock()
	if _, ok := t.lazy[p]; ok {
		delete(t.lazy, p)
	}
	t.lazyMu.Unlock()

	t.eagerMu.Lock()
	defer t.eagerMu.Unlock()
	_, ok := t.eager[p]
	if ok {
		return fmt.Errorf("peer %s already in eager peerset", p.String())
	}
	t.eager[p] = struct{}{}
	return nil
}

// moveToLazy moves the given peer to the lazy peerset
func (t *tracker) moveToLazy(p peer.ID) error {
	log.Debug("moving peer to lazy set", "peer", p.String())
	t.eagerMu.Lock()
	if _, ok := t.eager[p]; ok {
		delete(t.eager, p)
	}
	t.eagerMu.Unlock()

	t.lazyMu.Lock()
	defer t.lazyMu.Unlock()
	_, ok := t.lazy[p]
	if ok {
		return fmt.Errorf("peer %s already in lazy peerset", p.String())
	}
	t.lazy[p] = struct{}{}
	return nil
}
