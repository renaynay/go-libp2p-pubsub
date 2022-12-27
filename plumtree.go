package pubsub

import (
	libpeer "github.com/libp2p/go-libp2p/core/peer"
	"github.com/libp2p/go-libp2p/core/protocol"

	pb "github.com/libp2p/go-libp2p-pubsub/pb"
)

type Router struct {
	p       *PubSub
	tracker *tracker
}

func (r *Router) Protocols() []protocol.ID {
	//TODO implement me
	panic("implement me")
}

func (r *Router) Attach(sub *PubSub) {
	//TODO implement me
	panic("implement me")
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
	// eager peers get full message
	for p := range r.tracker.eager {
		r.p.peers[p] <- rpcWithMessages(msg.Message)
	}
	// send IHAVE of msg hash to lazy peers
	// TODO: schedule messages to be sent, instead of sending right away
	for p := range r.tracker.lazy {
		// send msg
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
