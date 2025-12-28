package config

import "os"

type Config struct {
	Port     string
	Redis    RedisConfig
	External ExternalAPIConfig
}

type RedisConfig struct {
	Addr string
}

type ExternalAPIConfig struct {
	BaseURL string
}

func Load() *Config {
	return &Config{
		Port: getEnv("PORT", "8080"),
		Redis: RedisConfig{
			Addr: getEnv("REDIS_ADDR", "localhost:6379"),
		},
		External: ExternalAPIConfig{
			BaseURL: getEnv("FOOD_API_BASE_URL", "https://world.openfoodfacts.org/api/v0"),
		},
	}
}

func getEnv(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}

