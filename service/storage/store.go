package storage

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

type DB struct {
	conn *pgxpool.Pool
}

func NewStore(conn *pgxpool.Pool) *DB {
	return &DB{
		conn: conn,
	}
}

func (db *DB) StoreCandles(ctx context.Context, c OHLC) error {
	sql := `INSERT INTO candlesticks (
	symbol, start_ts, open, close, high, low , volume)
	VALUES ($1, $2 , $3, $4, $5, $6, $7)`

	_, err := db.conn.Exec(ctx, sql,
		c.Symbol,
		c.StartTime,
		c.Open,
		c.Close,
		c.High,
		c.Low,
		c.Volume,
	)

	return err
}
