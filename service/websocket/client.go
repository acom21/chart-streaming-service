package websocket

import (
	"context"
	"encoding/json"
	"fmt"
	"net/url"
	"strconv"
	"time"

	"github.com/acom21/chart-streaming-service/service/aggregator"
	"github.com/gorilla/websocket"
	"github.com/shopspring/decimal"
	"go.uber.org/zap"
)

type Client struct {
	URL    string
	Dialer *websocket.Dialer
	Logger *zap.Logger
}

func NewClient(url string, dialer *websocket.Dialer, logger *zap.Logger) *Client {
	return &Client{
		URL:    url,
		Dialer: dialer,
		Logger: logger,
	}
}

func (c *Client) ConnectAndListen(ctx context.Context) <-chan aggregator.Event {
	v := url.Values{}
	v.Add("streams", "btcusdt@trade/ethusdt@trade/pepeusdt@trade")

	u, _ := url.Parse(c.URL + "?" + v.Encode())
	out := make(chan aggregator.Event)

	conn, _, err := c.Dialer.Dial(u.String(), nil)
	if err != nil {

		out <- aggregator.Event{
			Err: fmt.Errorf("websocket dial err %w", err),
		}
		return out
	}

	go func() {
		defer conn.Close()
		defer close(out)

		for {
			select {
			case <-ctx.Done():

				c.Logger.Info("WebSocket closed")
				return
			default:
				_, msg, err := conn.ReadMessage()
				if err != nil {
					c.Logger.Error("read error", zap.Error(err))
					out <- aggregator.Event{
						Err: fmt.Errorf("read error %w", err),
					}
					continue
				}

				var m Message
				if err := json.Unmarshal(msg, &m); err != nil {
					c.Logger.Error("json error:", zap.Error(err))

					out <- aggregator.Event{Err: fmt.Errorf("json error: %w", err)}
					continue
				}

				price, err := decimal.NewFromString(m.Data.Price)
				if err != nil {
					c.Logger.Error("failed parse price", zap.Error(err))

					out <- aggregator.Event{
						Err: fmt.Errorf("failed parse price %w", err),
					}
					continue
				}
				qty, err := strconv.ParseFloat(m.Data.Quantity, 64)
				if err != nil {
					c.Logger.Error("failed parse quantity", zap.Error(err))

					out <- aggregator.Event{
						Err: fmt.Errorf("failed parse quantity %w", err),
					}

					continue
				}

				out <- aggregator.Event{
					Trade: &aggregator.TradeData{
						Symbol:    m.Data.Symbol,
						Price:     price,
						Quantity:  qty,
						Timestamp: time.UnixMilli(m.Data.TradeTime),
					},
					Err: nil,
				}

			}
		}
	}()

	return out
}
