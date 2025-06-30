package config

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

type Config struct {
	LogLevel string
	DBURL    string
	GRPC     string
	App      AppConfig
	Dialer   DialerConfig
}

type AppConfig struct {
	BaseURL  string
	Interval time.Duration
}

type DialerConfig struct {
	HandshakeTimeout time.Duration
	KeepAlive        time.Duration
}

func NewConfig() (*Config, error) {
	port := os.Getenv("DB_PORT")
	if port == "" {
		port = "5432"
	}
	dbURL := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=%s",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		port,
		os.Getenv("DB_NAME"),
		os.Getenv("DB_SSLMODE"),
	)

	interval, err := time.ParseDuration(os.Getenv("APP_INTERVAL"))
	if err != nil {
		return nil, fmt.Errorf("invalid APP_INTERVAL: %w", err)
	}

	handshakeMs, err := strconv.Atoi(os.Getenv("DIALER_HANDSHAKE_TIMEOUT_MS"))
	if err != nil {
		return nil, fmt.Errorf("invalid handshake timeout: %w", err)
	}

	keepaliveSec, err := strconv.Atoi(os.Getenv("DIALER_KEEPALIVE_SECONDS"))
	if err != nil {
		return nil, fmt.Errorf("invalid keepalive: %w", err)
	}

	return &Config{
		LogLevel: os.Getenv("LOG_LEVEL"),
		DBURL:    dbURL,
		GRPC:     os.Getenv("GRPC_ADDRESS") + ":" + os.Getenv("GRPC_PORT"),
		App: AppConfig{
			BaseURL:  os.Getenv("APP_BASE_URL"),
			Interval: interval,
		},
		Dialer: DialerConfig{
			HandshakeTimeout: time.Duration(handshakeMs) * time.Millisecond,
			KeepAlive:        time.Duration(keepaliveSec) * time.Second,
		},
	}, nil
}
