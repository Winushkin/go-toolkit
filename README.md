# My Go Toolkit 🛠️

Собственный набор готовых модулей и оберток для быстрой инициализации Go-сервисов. Помогает избежать копирования шаблонного кода для логгера и базы данных из проекта в проект.

## Состав пакета
- `config` — быстрый сбор переменных окружения с дефолтными значениями.
- `logger` — преднастроенный структурированный JSON-логгер на базе стандарта `log/slog`.
- `postgres` — пул подключений к PostgreSQL на базе драйвера `pgx` с оптимальными настройками.

## Быстрый старт

### 1. Установка пакета

Подключите библиотеку к вашему новому Go-проекту:

```bash
go get github.com/Winushkin/go-toolkit
```


### 2. Пример использования

Добавьте этот код в `main.go` вашего нового сервиса:

```go
package main

import (
	"context"
	"fmt"

	"github.com/Winushkin/go-toolkit/config"
	"github.com/Winushkin/go-toolkit/logger"
	"github.com/Winushkin/go-toolkit/postgres"
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
	pgCfg := postgres.Config{
		Host: config.GetEnv("POSTGRES_HOST", "localhost"),
		Port: config.GetEnv("POSTGRES_PORT", "5432"),
		// ...
		MaxConns: config.GetEnv("POSTGRES_MAXCONNS", "1"),
	}

	// 3. Подключение к PostgreSQL
	pool, err := postgres.NewPool(ctx, pgCfg)
	if err != nil {
		panic(fmt.Errorf("failed to create postgres pool: %w", err))
	}

	log.Info("Успешное подключение к базе данных!")
}
```