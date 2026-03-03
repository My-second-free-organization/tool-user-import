package config

import "github.com/spf13/viper"

type Config struct {
	ServerPort     int    `mapstructure:"SERVER_PORT"`
	DatabaseURL    string `mapstructure:"DATABASE_URL"`
	RedisURL       string `mapstructure:"REDIS_URL"`
	JWTSecret      string `mapstructure:"JWT_SECRET"`
	CoreServiceURL string `mapstructure:"CORE_SERVICE_URL"`
	LogLevel       string `mapstructure:"LOG_LEVEL"`
	RateLimitRPS   int    `mapstructure:"RATE_LIMIT_RPS"`
}

func Load() (*Config, error) {
	viper.SetDefault("SERVER_PORT", 8080)
	viper.SetDefault("LOG_LEVEL", "info")
	viper.SetDefault("RATE_LIMIT_RPS", 100)
	viper.AutomaticEnv()
	var cfg Config
	viper.Unmarshal(&cfg)
	return &cfg, nil
}
