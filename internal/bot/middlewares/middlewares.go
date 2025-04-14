package middlewares

import (
	"context"
	"log"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

func LogMessage(next bot.HandlerFunc) bot.HandlerFunc {
	return func(ctx context.Context, b *bot.Bot, update *models.Update) {
		if update.Message != nil {
			log.Printf("Message from %v: %v", update.Message.From.FirstName, update.Message.Text)
		} else if update.CallbackQuery != nil {
			log.Printf("Callback from %v: %v", update.CallbackQuery.From.FirstName, update.CallbackQuery.Data)
		}
		next(ctx, b, update)
	}
}
