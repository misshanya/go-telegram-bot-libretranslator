package main

import (
	"context"
	//"log"
	"os"
	"os/signal"

	"github.com/go-telegram/bot"
	"github.com/misshanya/go-telegram-bot-libretranslator/internal/config"
	"github.com/misshanya/go-telegram-bot-libretranslator/internal/handlers"
)

func main() {
	cfg := config.GetConfig()

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	b, err := bot.New(cfg.TelegramToken)
	if err != nil {
		panic(err)
	}

	handlers.RegisterHandlers(b)

	b.Start(ctx)
}
