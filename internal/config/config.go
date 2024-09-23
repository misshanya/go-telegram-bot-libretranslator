package config

import (
	"os"
	"sync"

	"github.com/joho/godotenv"
)

type Config struct {
	TelegramToken     string
	LibreTranslateUrl string
}

var (
	config *Config
	once   sync.Once
)

func loadConfig() {
	err := godotenv.Load()
	if err != nil {
		config = &Config{}
		return
	}

	config = &Config{
		TelegramToken:     os.Getenv("BOT_TOKEN"),
		LibreTranslateUrl: os.Getenv("LIBRETRANSLATE_URL"),
	}
}

func GetConfig() *Config {
	once.Do(loadConfig)
	return config
}
