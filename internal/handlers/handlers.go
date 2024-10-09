package handlers

import (
	"context"
	"strings"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
	"github.com/misshanya/go-telegram-bot-libretranslator/internal/keyboard/reply"
	"github.com/misshanya/go-telegram-bot-libretranslator/internal/utils"
)

func RegisterHandlers(b *bot.Bot) {
	b.RegisterHandler(bot.HandlerTypeMessageText, "/start", bot.MatchTypeExact, startHandler)
	b.RegisterHandler(bot.HandlerTypeMessageText, "/translate", bot.MatchTypePrefix, translateHandler)
}

func startHandler(ctx context.Context, b *bot.Bot, update *models.Update) {
	_ = utils.RegisterUser(ctx, update.Message.From.ID)
	b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID:      update.Message.Chat.ID,
		Text:        "Привет! Я бот-переводчик. Написан на Go. Для перевода используется API LibreTranslate",
		ReplyMarkup: reply.InitReplyKeyboard(b),
	})
}

func translateHandler(ctx context.Context, b *bot.Bot, update *models.Update) {
	textToTranslate := strings.TrimSpace(strings.TrimPrefix(update.Message.Text, "/translate"))

	langFrom := "ru"
	var langTo string

	if utils.IsAutoDetect(ctx, update.Message.From.ID) {
		langFrom = utils.DetectLanguage(textToTranslate)
	}

	if langFrom == "ru" {
		langTo = "en"
	} else {
		langTo = "ru"
	}

	translatedText := utils.Translate(textToTranslate, langFrom, langTo)

	b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID: update.Message.Chat.ID,
		Text:   translatedText,
	})
}
