package config

import (
	dotenv "github.com/joho/godotenv"
	"os"
)

type Config struct {
	BotToken   string
	DbHost     string
	DbPort     string
	DbUser     string
	DbPassword string
	DbName     string
}

func MustLoad() *Config {
	if err := dotenv.Load(); err != nil {
		panic("Fuck my environment")
	}

	return &Config{
		BotToken:   os.Getenv("BOT_TOKEN"),
		DbHost:     os.Getenv("DB_HOST"),
		DbPort:     os.Getenv("DB_PORT"),
		DbUser:     os.Getenv("DB_USER"),
		DbPassword: os.Getenv("DB_PASSWORD"),
		DbName:     os.Getenv("DB_NAME"),
	}
}
