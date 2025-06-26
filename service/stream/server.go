package stream

import pb "github.com/acom21/chart-streaming-service/service/stream/proto/tick/pb/tick"

type Server struct {
	pb.UnimplementedStreamingServiceServer
	SubMgr *SubscriptionMgr
}

func (s *Server) StreamTick(req *pb.StreamRequest, stream pb.StreamingService_StreamTickServer) error {
	sub := s.SubMgr.Subscribe(req.Symbols)
	defer s.SubMgr.Unsubscribe(sub)

	for tick := range sub.Ch {
		protoTick := &pb.Tick{
			Symbol:   tick.Symbol,
			Ts:       tick.Ts,
			TradeId:  tick.TradeId,
			Quantity: tick.Quantity,
			Price:    tick.Price,
		}

		if err := stream.Send(protoTick); err != nil {
			return err
		}
	}

	return nil
}
