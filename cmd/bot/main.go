package main

import (
	"context"
	"flag"
	"os"
	"os/signal"

	"github.com/go-telegram/bot"
	"github.com/misshanya/go-telegram-bot-libretranslator/internal/config"
	"github.com/misshanya/go-telegram-bot-libretranslator/internal/handlers"
	"github.com/misshanya/go-telegram-bot-libretranslator/internal/middlewares"
)

func main() {
	logEnabled := flag.Bool("mlog", false, "Enable logging with middlewares")
	flag.Parse()

	cfg := config.GetConfig()
	_ = config.GetDB()

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	var opts []bot.Option

	if *logEnabled {
		opts = []bot.Option{
			bot.WithMiddlewares(middlewares.LogMessage),
		}
	}

	b, err := bot.New(cfg.TelegramToken, opts...)
	if err != nil {
		panic(err)
	}

	handlers.RegisterHandlers(b)

	b.Start(ctx)
}
