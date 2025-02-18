package config

import (
	"database/sql"
	"log"
	"os"
	"sync"

	"github.com/joho/godotenv"
	_ "github.com/mattn/go-sqlite3"
	"github.com/misshanya/go-telegram-bot-libretranslator/internal/db/users"
)

type Config struct {
	TelegramToken     string
	LibreTranslateUrl string
}

var (
	config  *Config
	dbConn  *sql.DB
	queries *users.Queries
	onceCfg sync.Once
	onceDB  sync.Once
)

// loadConfig loads environment variables from a .env file and initializes the config.
// If loading fails, it initializes an empty Config.
func loadConfig() {
	err := godotenv.Load()
	if err != nil {
		log.Println(".env file not found, trying to use environment variables...")
	}

	botToken := os.Getenv("BOT_TOKEN")
	if botToken == "" {
		log.Fatalln("missing BOT_TOKEN env var")
	}

	libreTranslateUrl := os.Getenv("LIBRETRANSLATE_URL")
	if libreTranslateUrl == "" {
		log.Fatalln("missing LIBRETRANSLATE_URL env var")
	}

	config = &Config{
		TelegramToken:     botToken,
		LibreTranslateUrl: libreTranslateUrl,
	}
}

// initDB initializes the database connection and creates the users table if it doesn't exist.
// It logs a fatal error if the connection or table creation fails.
func initDB() {
	var err error
	dbConn, err = sql.Open("sqlite3", "./bot.db")
	if err != nil {
		log.Fatalln("Failed to connect to database:", err)
	}
	queries = users.New(dbConn)
}

// GetConfig returns the application configuration, loading it only once.
func GetConfig() *Config {
	onceCfg.Do(loadConfig)
	return config
}

// GetDB returns the database queries instance, initializing the database only once.
func GetDB() *users.Queries {
	onceDB.Do(initDB)
	return queries
}
