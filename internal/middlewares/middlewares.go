package middlewares

import (
	"context"
	"log"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

func LogMessage(next bot.HandlerFunc) bot.HandlerFunc {
	return func(ctx context.Context, b *bot.Bot, update *models.Update) {
		log.Printf("Message from %v: %v", update.Message.From.FirstName, update.Message.Text)
		next(ctx, b, update)
	}
}
