package aggregator

import (
	"context"
	"time"

	"github.com/acom21/chart-streaming-service/service/storage"
	"go.uber.org/zap"
)

// TradeStreamer represent websocket stream inerface.
type TradeStreamer interface {
	ConnectAndListen(ctx context.Context) <-chan Event
}

// Store represent db inerface.
type Store interface {
	StoreCandles(ctx context.Context, c storage.OHLC) error
}

// Aggregator represents main service  struct.
type Aggregator struct {
	current  map[string]storage.OHLC
	interval time.Duration
	log      *zap.Logger
	ts       TradeStreamer
	store    Store
}

func NewAggregator(interval time.Duration, ts TradeStreamer, store Store, log *zap.Logger) Aggregator {
	return Aggregator{
		current:  make(map[string]storage.OHLC),
		interval: interval,
		log:      log,
		ts:       ts,
		store:    store,
	}
}

func (a *Aggregator) Aggregate(ctx context.Context) {

	a.log.Info("service  started")
	eventChan := a.ts.ConnectAndListen(ctx)
	ticker := time.NewTicker(a.interval)
	defer ticker.Stop()

	for {
		select {
		case ev, ok := <-eventChan:
			if !ok {
				return
			}
			if ev.Err != nil {
				a.log.Error("stream error", zap.Error(ev.Err))
				continue
			}
			a.createOHLC(ctx, *ev.Trade)

		case now := <-ticker.C:
			a.finalizeInterval(ctx, now)

		case <-ctx.Done():
			a.log.Info("service succesefully finished")
			return
		}
	}

}

func (a *Aggregator) createOHLC(ctx context.Context, t TradeData) {
	start := t.Timestamp.Truncate(a.interval)
	key := t.Symbol

	c, exists := a.current[key]
	if !exists || c.StartTime.Before(start) {
		if exists {
			a.storeAndReset(ctx, key, c)
		}
		a.current[key] = storage.OHLC{
			Symbol:    t.Symbol,
			StartTime: start,
			Open:      t.Price,
			High:      t.Price,
			Low:       t.Price,
			Close:     t.Price,
			Volume:    t.Quantity,
		}
		return
	}

	// внутри того же интервала
	if t.Price.GreaterThan(c.High) {
		c.High = t.Price
	}
	if t.Price.LessThan(c.Low) {
		c.Low = t.Price
	}
	c.Close = t.Price
	c.Volume += t.Quantity
	a.current[key] = c
}

func (a *Aggregator) finalizeInterval(ctx context.Context, now time.Time) {
	boundary := now.Truncate(a.interval)
	for key, c := range a.current {
		if c.StartTime.Before(boundary) {
			a.storeAndReset(ctx, key, c)
		}
	}
}

func (a *Aggregator) storeAndReset(ctx context.Context, key string, c storage.OHLC) {
	if err := a.store.StoreCandles(ctx, c); err != nil {
		a.log.Error("failed to store candle", zap.Error(err), zap.String("symbol", c.Symbol), zap.Time("start", c.StartTime))
		return
	}
	delete(a.current, key)

	a.log.Info("candle inserted to db", zap.Any("ohlc", c))
}
