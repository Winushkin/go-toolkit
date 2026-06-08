// Package redis содержит базовое поключение к Redis DB
package redis

import (
	"context"
	"fmt"
	"net"

	"github.com/redis/go-redis/v9"
)

type Config struct {
	Host     string `env:"HOST"`
	Port     string `env:"PORT"`
	Password string `env:"PASSWORD"`
	DB       int    `env:"DB"`
}

func NewRedisClient(ctx context.Context, cfg Config) (*redis.Client, error) {
	rdb := redis.NewClient(&redis.Options{
		Addr:     net.JoinHostPort(cfg.Host, cfg.Port),
		Password: cfg.Password,
		DB:       cfg.DB,
	})

	if err := rdb.Ping(ctx).Err(); err != nil {
		return nil, fmt.Errorf("failed to connect to Redis: %v", err)
	}

	return rdb, nil
}
