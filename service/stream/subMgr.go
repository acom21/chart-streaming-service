package stream

import (
	"slices"
	"sync"

	pb "github.com/acom21/chart-streaming-service/service/stream/proto/tick/pb/tick"
)

type SubscriptionMgr struct {
	mu   sync.RWMutex
	subs map[string][]chan *pb.Tick
}

func NewSubManager() *SubscriptionMgr {
	return &SubscriptionMgr{
		subs: make(map[string][]chan *pb.Tick),
	}
}

func (m *SubscriptionMgr) Subscribe(symbol string, ch chan *pb.Tick) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.subs[symbol] = append(m.subs[symbol], ch)
}

func (m *SubscriptionMgr) Unsubscribe(symbol string, ch chan *pb.Tick) {
	m.mu.Lock()
	defer m.mu.Unlock()
	chs := m.subs[symbol]
	for i, c := range chs {
		if c == ch {
			m.subs[symbol] = slices.Delete(chs, i, i+1)
			break
		}
	}
}

func (m *SubscriptionMgr) Broadcast(symbol string, tick *pb.Tick) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	for _, ch := range m.subs[symbol] {
		select {
		case ch <- tick:
		default:
		}

	}
}
