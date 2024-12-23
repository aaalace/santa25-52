package config

import (
	dotenv "github.com/joho/godotenv"
	"os"
)

type Config struct {
	BotToken  string
	RedisHost string
	RedisPort string
}

func MustLoad() *Config {
	if err := dotenv.Load(); err != nil {
		panic("Fuck my environment")
	}

	return &Config{
		BotToken:  os.Getenv("BOT_TOKEN"),
		RedisHost: os.Getenv("REDIS_HOST"),
		RedisPort: os.Getenv("REDIS_PORT"),
	}
}
