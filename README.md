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
```

#### Если ваши переменные называются стандартно:
```
POSTGRES_HOST
POSTGRES_PORT
POSTGRES_USER
POSTGRES_PASSWORD
POSTGRES_DB
POSTGRES_MIN_CONNS
POSTGRES_MAX_CONNS

SERVER_PORT
SERVER_HOST
DOMAIN_NAME
```
то можно использовать функцию, которая автоматически подтянет ваши изменения:
```go
	// 2. Инициализация конфигурации PostgreSQL
	cfg := config.NewAppConfig()
	pgCfg := cfg.Postgres
```
Если же ваши переменные называются иначе, то придется использовать функция GetEnv и считывать каждую переменную отдельно: 
```go
	// 2. Инициализация конфигурации PostgreSQL
	pgCfg := postgres.Config{
		Host: config.GetEnv("POSTGRES_HOST", "localhost"),
		Port: config.GetEnv("POSTGRES_PORT", "5432"),
		// ...
		MaxConns: config.GetEnv("POSTGRES_MAXCONNS", "1"),
	}
```

И далее использовать ваш конфиг
```go
	// 3. Подключение к PostgreSQL
	pool, err := postgres.NewPool(ctx, pgCfg)
	if err != nil {
		log.Error(ctx, err, "failed to create postgres pool")
	}

	log.Info(ctx, "Успешное подключение к базе данных!", zap.String("Port", pgCfg.Port))
	pool.Exec(ctx, "query")
}
```