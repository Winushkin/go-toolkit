package main

import (
	"context"
	"fmt"

	"github.com/Winushkin/go-toolkit/config"
	"github.com/Winushkin/go-toolkit/logger"
	"github.com/Winushkin/go-toolkit/postgres"
	"go.uber.org/zap"
)

const devMode = true

func main() {
	// 1. Инициализация логгера
	ctx, err := logger.NewLoggerContext(context.Background(), devMode)
	if err != nil {
		panic(fmt.Errorf("failed to create logger context: %w", err))
	}

	log, ok := logger.GetLoggerFromCtx(ctx)
	if !ok {
		panic("logger not found in context")
	}
	// 2. Инициализация конфигурации PostgreSQL
	cfg := config.NewAppConfig()
	pgCfg := cfg.Postgres

	// 3. Подключение к PostgreSQL
	pool, err := postgres.NewPool(ctx, pgCfg)
	if err != nil {
		panic(fmt.Errorf("failed to create postgres pool: %w", err))
	}

	log.Info(ctx, "Успешное подключение к базе данных!", zap.String("Port", pgCfg.Port))
	pool.Exec(ctx, "query")
}
