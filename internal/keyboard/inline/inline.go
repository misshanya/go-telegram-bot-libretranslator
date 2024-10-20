package inline

import (
	"context"
	"fmt"
	"log"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
	"github.com/go-telegram/ui/keyboard/inline"
	"github.com/misshanya/go-telegram-bot-libretranslator/internal/utils"
)

func InitInlineKeyboard(ctx context.Context, b *bot.Bot, update *models.Update) *inline.Keyboard {
	return createKeyboard(ctx, b, update.Message.From.ID)
}

func onInlineKeyboardSelect(ctx context.Context, b *bot.Bot, mes models.MaybeInaccessibleMessage, data []byte) {
	switch string(data) {
	case "lang-autodetect":
		utils.ChangeAutoDetect(ctx, mes.Message.Chat.ID)
		updateKeyboard(ctx, b, &mes)

	case "lang-from":
		// TODO: lang-from callback handling

	case "lang-to":
		// TODO: lang-to callback handling
	}
}

func createKeyboard(ctx context.Context, b *bot.Bot, uid int64) *inline.Keyboard {
	autoDetectText := fmt.Sprintf("Автоопределение языка: %v", getAutoDetectChar(ctx, uid))
	langFromText := fmt.Sprintf("Переводить с: %v", "язык")
	langToText := fmt.Sprintf("Переводить на: %v", "язык")

	return inline.New(b, inline.NoDeleteAfterClick()).
		Row().
		Button(autoDetectText, []byte("lang-autodetect"), onInlineKeyboardSelect).
		Row().
		Button(langFromText, []byte("lang-from"), onInlineKeyboardSelect).
		Button(langToText, []byte("lang-to"), onInlineKeyboardSelect)
}

func updateKeyboard(ctx context.Context, b *bot.Bot, mes *models.MaybeInaccessibleMessage) {
	newKb := createKeyboard(ctx, b, mes.Message.Chat.ID)

	_, err := b.EditMessageReplyMarkup(ctx, &bot.EditMessageReplyMarkupParams{
		ChatID:      mes.Message.Chat.ID,
		MessageID:   mes.Message.ID,
		ReplyMarkup: newKb,
	})
	if err != nil {
		log.Println("Error when updating keyboard:", err)
	}
}

func getAutoDetectChar(ctx context.Context, uid int64) string {
	var autoDetectChar string
	if utils.IsAutoDetect(ctx, uid) {
		autoDetectChar = "✅"
	} else {
		autoDetectChar = "❎"
	}
	return autoDetectChar
}
