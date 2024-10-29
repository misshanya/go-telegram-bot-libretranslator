package handlers

import (
	"context"
	"log"
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

	var sourceLang string = "ru"
	var targetLang string
	var err error

	if utils.IsAutoDetect(ctx, update.Message.From.ID) {
		sourceLang, err = utils.DetectLanguage(textToTranslate)
		if err != nil {
			log.Println(err)
			b.SendMessage(ctx, &bot.SendMessageParams{
				ChatID: update.Message.Chat.ID,
				Text:   "Возникла ошибка при определении языка",
			})
			return
		}
	}

	if sourceLang == "ru" {
		targetLang = "en"
	} else {
		targetLang = "ru"
	}

	translatedText, err := utils.Translate(textToTranslate, sourceLang, targetLang)
	if err != nil {
		log.Println(err)
		b.SendMessage(ctx, &bot.SendMessageParams{
			ChatID: update.Message.Chat.ID,
			Text:   "Возникла ошибка при переводе текста",
		})
		return
	}

	b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID: update.Message.Chat.ID,
		Text:   translatedText,
	})
}
