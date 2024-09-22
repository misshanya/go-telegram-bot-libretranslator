package config

import (
	"github.com/joho/godotenv"
	"os"
)

type Config struct {
	TelegramToken     string
	LibreTranslateUrl string
}

func LoadConfig() (Config, error) {
	err := godotenv.Load()
	if err != nil {
		return Config{}, err
	}

	return Config{
		TelegramToken:     os.Getenv("BOT_TOKEN"),
		LibreTranslateUrl: os.Getenv("LIBRETRANSLATE_URL"),
	}, nil
}
