package config

import (
	"os"
	"strconv"
)

type AppConfig struct {
	GrpcPort         int
	PgConUrl         string
	PgMigrationsPath string
	LogLevel         string
}

func LoadConfigFromEnv() AppConfig {
	return AppConfig{
		GrpcPort:         getEnvInt("GRPC_PORT", 9000),
		PgConUrl:         getEnv("PG_CONNECTION_URL", "postgres://postgres:super_secret_123@localhost/incrementor?sslmode=disable"),
		PgMigrationsPath: getEnv("PG_MIGRATIONS_PATH", "file://../../internal/database/migrations"),
		LogLevel:         getEnv("LOG_LEVEL", "debug"),
	}
}

func getEnv(key string, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}

func getEnvInt(key string, defaultValue int) int {
	valueStr := os.Getenv(key)
	if value, err := strconv.Atoi(valueStr); err == nil {
		return value
	}
	return defaultValue
}
