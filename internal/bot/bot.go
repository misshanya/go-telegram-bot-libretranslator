package bot

import (
	"context"
	"github.com/go-telegram/bot"
	"github.com/misshanya/go-telegram-bot-libretranslator/internal/bot/handlers"
	"github.com/misshanya/go-telegram-bot-libretranslator/internal/bot/middlewares"
	"github.com/misshanya/go-telegram-bot-libretranslator/internal/bot/ui/keyboard/inline"
	"github.com/misshanya/go-telegram-bot-libretranslator/internal/bot/ui/keyboard/reply"
	"github.com/misshanya/go-telegram-bot-libretranslator/internal/config"
	"github.com/misshanya/go-telegram-bot-libretranslator/internal/repository"
	"github.com/misshanya/go-telegram-bot-libretranslator/internal/service"
	"log"
)

func Start(logEnabled *bool, cfg *config.Config, ctx context.Context) {
	// Init queries
	queries := config.GetDB()

	// Init repo
	botRepo := repository.NewRepository(queries)

	// Init service
	botService := service.NewService(botRepo)

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

	// Init handlers
	botInlineKb := inline.NewInlineKeyboard(botService)
	botReplyKb := reply.NewReplyKeyboard(botInlineKb)
	botHandler := handlers.NewHandler(b, botReplyKb, botService)
	botHandler.RegisterHandlers()

	log.Println("Starting bot")

	b.Start(ctx)
}
