package stream

import (
	"sync"

	pb "github.com/acom21/chart-streaming-service/service/stream/proto/tick/pb/tick"
)

type Subscription struct {
	Symbols map[string]struct{}
	Ch      chan pb.Tick
}

type SubscriptionMgr struct {
	mu   sync.RWMutex
	subs map[*Subscription]struct{}
}

func (m *SubscriptionMgr) Subscribe(symbols []string) *Subscription {
	symMap := make(map[string]struct{}, len(symbols))
	for _, s := range symbols {
		symMap[s] = struct{}{}
	}

	sub := &Subscription{
		Symbols: symMap,
		Ch:      make(chan pb.Tick),
	}

	m.mu.Lock()
	m.subs[sub] = struct{}{}
	m.mu.Unlock()

	return sub
}

func (m *SubscriptionMgr) Unsubscribe(sub *Subscription) {
	m.mu.Lock()
	delete(m.subs, sub)
	m.mu.Unlock()

	close(sub.Ch)
}

func (m *SubscriptionMgr) Publish(t pb.Tick) {
	m.mu.RLock()
	defer m.mu.RUnlock()

	for sub := range m.subs {
		if _, ok := sub.Symbols[t.Symbol]; ok {
			select {
			case sub.Ch <- t:
			default:

			}
		}
	}
}
