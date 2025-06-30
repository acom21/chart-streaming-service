package main

import (
	"context"
	"fmt"
	"net"
	"time"

	"github.com/acom21/chart-streaming-service/pkg/config"
	"github.com/acom21/chart-streaming-service/service/aggregator"
	"github.com/acom21/chart-streaming-service/service/storage"
	"github.com/acom21/chart-streaming-service/service/stream"
	"github.com/acom21/chart-streaming-service/service/websocket"

	pb "github.com/acom21/chart-streaming-service/service/stream/proto/tick/pb/tick"

	ws "github.com/gorilla/websocket"
	pgxdecimal "github.com/jackc/pgx-shopspring-decimal"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func run(ctx context.Context) error {
	cfg, err := config.NewConfig()
	if err != nil {
		return fmt.Errorf("failed to load config: %w", err)
	}

	logger, err := initLogger(cfg.LogLevel)
	if err != nil {
		return fmt.Errorf("failed to init logger: %w", err)
	}
	defer logger.Sync()

	conn, err := initPostgres(ctx, cfg.DBURL)
	if err != nil {
		return fmt.Errorf("init db %w", err)
	}

	subMgr := stream.NewSubManager()
	grpcService := stream.NewGRPCServer(subMgr)
	lis, _ := net.Listen("tcp", cfg.GRPC)
	grpcSrv := grpc.NewServer()

	pb.RegisterStreamingServiceServer(grpcSrv, grpcService)
	go grpcSrv.Serve(lis)

	reflection.Register(grpcSrv)

	store := storage.NewStore(conn)
	websocket := websocket.NewClient(
		cfg.App.BaseURL,
		ws.DefaultDialer,
		logger,
	)

	aggr := aggregator.NewAggregator(cfg.App.Interval, websocket, grpcService, store, logger)

	aggr.Aggregate(ctx)

	return nil
}

func initLogger(logLevel string) (*zap.Logger, error) {
	var lvl zap.AtomicLevel
	switch logLevel {
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

func initPostgres(ctx context.Context, dbURL string) (*pgxpool.Pool, error) {
	cfg, err := pgxpool.ParseConfig(dbURL)
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
