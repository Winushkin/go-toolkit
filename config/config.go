// Package config содержит функцию для сичтывания переменных окружения
package config

import "os"

// GetEnv помогает быстро читать переменные с дефолтными значениями в ваших проектах
func GetEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}
