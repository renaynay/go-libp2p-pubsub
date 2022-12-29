package pubsub

import (
	libpeer "github.com/libp2p/go-libp2p/core/peer"
	"github.com/libp2p/go-libp2p/core/protocol"

	pb "github.com/libp2p/go-libp2p-pubsub/pb"
)

// Router is an implementation of a Plumtree router.
/*
TODO:
  - Implement message cache (@distractedm1nd)
  - Implement timed publishing for lazies (@distractedm1nd)
  - Implement timer on receiving IHAVE -> sending IWANT/GRAFT if not already received by direct peer (the peer that has you as an eager)
  - Implement listener loop for messages (checking for dups and sending PRUNEs) - receive handlers (@renaynay)
  - Implement Optimization procedure for tree cleanup (@distractedm1nd) - NICE TO HAVE RN
  - implement NeighborUp/NeighborDown (@renaynay)
*/
type Router struct {
	p       *PubSub
	tracker *tracker
}

func (r *Router) Protocols() []protocol.ID {
	//TODO implement me
	panic("implement me")
}

func (r *Router) Attach(sub *PubSub) {
	r.p = sub

	r.tracker.start()
}

func (r *Router) AddPeer(id libpeer.ID, id2 protocol.ID) {
	// neighborUp
	//TODO implement me
	panic("implement me")
}

func (r *Router) RemovePeer(id libpeer.ID) {
	// neighborDown
	//TODO implement me
	panic("implement me")
}

func (r *Router) EnoughPeers(topic string, suggested int) bool {
	//TODO implement me
	panic("implement me")
}

func (r *Router) AcceptFrom(id libpeer.ID) AcceptStatus {
	//TODO implement me
	panic("implement me")
}

func (r *Router) HandleRPC(rpc *RPC) {
	//TODO implement me
	panic("implement me")
}

func (r *Router) Publish(msg *Message) {
	// TODO: increase round parameter in message
	// eager peers get full message
	for p := range r.tracker.eager {
		r.p.peers[p] <- rpcWithMessages(msg.Message)
	}
	// send IHAVE of msg hash to lazy peers
	ihave := &RPC{
		RPC: pb.RPC{
			Control: &pb.ControlMessage{
				Ihave: []*pb.ControlIHave{{
					TopicID:    msg.Topic,
					MessageIDs: []string{msg.ID},
				}},
			},
		},
	}
	for p := range r.tracker.lazy {
		// TODO: schedule messages to be sent, instead of sending right away
		r.p.peers[p] <- ihave
	}
}

func (r *Router) Join(topic string) {
	//TODO implement me
	panic("implement me")
}

func (r *Router) Leave(topic string) {
	//TODO implement me
	panic("implement me")
}
