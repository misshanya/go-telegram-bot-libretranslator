package main

import (
	"context"
	"flag"
	"github.com/misshanya/go-telegram-bot-libretranslator/internal/bot"
	"os"
	"os/signal"

	"github.com/misshanya/go-telegram-bot-libretranslator/internal/config"
)

func main() {
	logEnabled := flag.Bool("mlog", false, "Enable logging with middlewares")
	flag.Parse()

	cfg := config.GetConfig()
	_ = config.GetDB()

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	bot.Start(logEnabled, cfg, ctx)
}
