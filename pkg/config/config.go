package config

import (
	"fmt"
	"os"
	"time"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Log    LogConfig    `yaml:"log"`
	DB     DBConfig     `yaml:"db"`
	Dialer DialerConfig `yaml:"dialer"`
	App    App          `yaml:"app"`
}

type LogConfig struct {
	Level string `yaml:"level"`
}

type DBConfig struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	DBName   string `yaml:"dbname"`
	SSLMode  string `yaml:"sslmode"`
}

type DialerConfig struct {
	HandshakeTimeoutMs int `yaml:"handshake_timeout_ms"`
	KeepAliveSeconds   int `yaml:"keepalive_seconds"`
}

type App struct {
	BaseURL  string        `yaml:"base_url"`
	Interval time.Duration `yaml:"interval"`
}

func NewConfig(filename string) (*Config, error) {
	var cfg Config

	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("read file: %w", err)
	}

	if err = yaml.Unmarshal(data, &cfg); err != nil {
		return nil, fmt.Errorf("parse YAML: %w", err)
	}

	return &cfg, nil
}
