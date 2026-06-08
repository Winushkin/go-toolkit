// Package config содержит функции для считывания конфигурационны переменных для приложения.
package config

import (
	"os"

	"github.com/Winushkin/go-toolkit/postgres"
	"github.com/Winushkin/go-toolkit/redis"
	"github.com/ilyakaznacheev/cleanenv"
)

// GetEnv помогает быстро читать переменные с дефолтными значениями в ваших проектах
func GetEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}

// AppConfig содержит конфигурационные переменные для приложения.
type AppConfig struct {
	Postgres   postgres.Config `env-prefix:"POSTGRES_"`
	Redis      redis.Config    `env-prefix:"REDIS_"`
	Port       string          `env:"SERVER_PORT"`
	Host       string          `env:"SERVER_HOST"`
	DomainName string          `env:"DOMAIN_NAME"`
}

// NewAppConfig считывает переменные окружения и возвращает структуру AppConfig.
func NewAppConfig() *AppConfig {
	var cfg AppConfig
	if err := cleanenv.ReadEnv(&cfg); err != nil {
		return nil
	}
	return &cfg
}
