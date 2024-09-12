package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	GRPCPort    string
	PostgresDSN string
	RedisAddr   string
	JWTSecret   string
}

func LoadConfig() *Config {
	viper.SetConfigFile(".env")
	viper.ReadInConfig()

	return &Config{
		GRPCPort:    viper.GetString("GRPC_PORT"),
		PostgresDSN: viper.GetString("POSTGRES_DSN"),
		RedisAddr:   viper.GetString("REDIS_ADDR"),
		JWTSecret:   viper.GetString("JWT_SECRET"),
	}
}
