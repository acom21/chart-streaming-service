package main

import (
	"context"
	"fmt"
	"time"

	"github.com/acom21/chart-streaming-service/pkg/config"
	"github.com/acom21/chart-streaming-service/service/aggregator"
	"github.com/acom21/chart-streaming-service/service/storage"
	pb "github.com/acom21/chart-streaming-service/service/stream/proto/tick/pb/tick"
	"github.com/acom21/chart-streaming-service/service/websocket"
	ws "github.com/gorilla/websocket"
	pgxdecimal "github.com/jackc/pgx-shopspring-decimal"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

func run(ctx context.Context, filePath string) error {
	cfg, err := config.NewConfig(filePath)
	if err != nil {
		return fmt.Errorf("failed to load config: %w", err)
	}

	logger, err := initLogger(cfg.Log)
	if err != nil {
		return fmt.Errorf("failed to init logger: %w", err)
	}
	defer logger.Sync()

	conn, err := initPostgres(ctx, cfg.DB)
	if err != nil {
		return fmt.Errorf("init db %w", err)
	}

	grepcSrv := grpc.NewServer()

	pb.RegisterStreamingServiceServer(srv, &MyTickServer{})

	store := storage.NewStore(conn)
	streamer := websocket.NewClient(
		cfg.App.BaseURL,
		ws.DefaultDialer,
		logger,
	)

	aggr := aggregator.NewAggregator(cfg.App.Interval, streamer, store, logger)

	aggr.Aggregate(ctx)

	return nil
}

func initLogger(cfg config.LogConfig) (*zap.Logger, error) {
	var lvl zap.AtomicLevel
	switch cfg.Level {
	case "debug":
		lvl = zap.NewAtomicLevelAt(zap.DebugLevel)
	case "info":
		lvl = zap.NewAtomicLevelAt(zap.InfoLevel)
	case "warn":
		lvl = zap.NewAtomicLevelAt(zap.WarnLevel)
	case "error":
		lvl = zap.NewAtomicLevelAt(zap.ErrorLevel)
	default:
		lvl = zap.NewAtomicLevelAt(zap.InfoLevel)
	}

	cfgZap := zap.Config{
		Level:         lvl,
		Development:   true,
		Encoding:      "console",
		OutputPaths:   []string{"stdout"},
		EncoderConfig: zap.NewDevelopmentEncoderConfig(),
	}
	return cfgZap.Build()
}

func initPostgres(ctx context.Context, dbCfg config.DBConfig) (*pgxpool.Pool, error) {
	connStr := fmt.Sprintf(
		"postgres://%s:%s@%s:%d/%s?sslmode=%s",
		dbCfg.User, dbCfg.Password, dbCfg.Host, dbCfg.Port, dbCfg.DBName, dbCfg.SSLMode,
	)
	cfg, err := pgxpool.ParseConfig(connStr)
	if err != nil {
		return nil, err
	}

	cfg.AfterConnect = func(ctx context.Context, c *pgx.Conn) error {

		pgxdecimal.Register(c.TypeMap())
		return nil
	}

	pool, err := pgxpool.NewWithConfig(ctx, cfg)
	if err != nil {
		return nil, err
	}

	ctxPing, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	if err := pool.Ping(ctxPing); err != nil {
		pool.Close()
		return nil, err
	}
	return pool, nil
}
