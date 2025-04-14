package handlers

import (
	"context"
	"fmt"
	"github.com/misshanya/go-telegram-bot-libretranslator/internal/bot/ui/keyboard/reply"
	"github.com/misshanya/go-telegram-bot-libretranslator/internal/service"
	"strings"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

type Handler struct {
	bot     *bot.Bot
	replyKb *reply.ReplyKeyboard
	service service.Service
}

func NewHandler(bot *bot.Bot, rkb *reply.ReplyKeyboard, service service.Service) *Handler {
	return &Handler{
		bot:     bot,
		replyKb: rkb,
		service: service,
	}
}

func (h *Handler) RegisterHandlers() {
	h.bot.RegisterHandler(bot.HandlerTypeMessageText, "/start", bot.MatchTypeExact, h.startHandler)
	h.bot.RegisterHandler(bot.HandlerTypeMessageText, "/translate", bot.MatchTypePrefix, h.translateHandler)
}

func (h *Handler) startHandler(ctx context.Context, b *bot.Bot, update *models.Update) {
	_ = h.service.RegisterUser(ctx, update.Message.From.ID)
	b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID:      update.Message.Chat.ID,
		Text:        "Привет! Я бот-переводчик. Написан на Go. Для перевода используется API LibreTranslate",
		ReplyMarkup: h.replyKb.InitReplyKeyboard(b),
	})
}

func (h *Handler) translateHandler(ctx context.Context, b *bot.Bot, update *models.Update) {
	textToTranslate := strings.TrimSpace(strings.TrimPrefix(update.Message.Text, "/translate"))

	translatedText, err := h.service.Translate(ctx, textToTranslate, update.Message.From.ID)
	if err != nil {
		b.SendMessage(ctx, &bot.SendMessageParams{
			ChatID: update.Message.Chat.ID,
			Text:   fmt.Sprint("Произошла ошибка при переводе"),
		})
	}

	b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID: update.Message.Chat.ID,
		Text:   translatedText,
	})
}
