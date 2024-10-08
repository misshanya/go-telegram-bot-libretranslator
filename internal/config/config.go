package config

import (
	"context"
	"database/sql"
	_ "embed"
	"log"
	"os"
	"sync"

	"github.com/joho/godotenv"
	_ "github.com/mattn/go-sqlite3"
	"github.com/misshanya/go-telegram-bot-libretranslator/internal/db/users"
	"github.com/misshanya/go-telegram-bot-libretranslator/sql/sqlUsers"
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

func initDB() {
	var err error
	ctx := context.Background()
	dbConn, err = sql.Open("sqlite3", "./bot.db")
	if err != nil {
		log.Fatalln("Failed to connect to database:", err)
	}
	if _, err := dbConn.ExecContext(ctx, sqlUsers.GetSchema()); err != nil {
		log.Fatalln("Failed to create table users:", err)
	}
	queries = users.New(dbConn)
}

func GetConfig() *Config {
	onceCfg.Do(loadConfig)
	return config
}

func GetDB() *users.Queries {
	onceDB.Do(initDB)
	return queries
}
