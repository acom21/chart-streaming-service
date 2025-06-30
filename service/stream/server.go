package stream

import (
	pb "github.com/acom21/chart-streaming-service/service/stream/proto/tick/pb/tick"
)

type Server struct {
	pb.UnimplementedStreamingServiceServer
	SubMgr *SubscriptionMgr
}

func NewGRPCServer(subMgr *SubscriptionMgr) *Server {
	return &Server{SubMgr: subMgr}
}

func (s *Server) StreamTick(req *pb.StreamRequest, stream pb.StreamingService_StreamTickServer) error {
	ctx := stream.Context()

	tickCh := make(chan *pb.Tick, 100)

	for _, symbol := range req.Symbols {
		s.SubMgr.Subscribe(symbol, tickCh)
		defer s.SubMgr.Unsubscribe(symbol, tickCh)
	}

	for {
		select {
		case <-ctx.Done():
			return ctx.Err()

		case tick := <-tickCh:
			if err := stream.Send(tick); err != nil {
				return err
			}
		}
	}
}

func (g *Server) Broadcast(symbol string, tick *pb.Tick) {
	g.SubMgr.Broadcast(symbol, tick)
}
