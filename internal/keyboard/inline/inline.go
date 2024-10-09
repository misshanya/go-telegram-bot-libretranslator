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
	kb := inline.New(b, inline.NoDeleteAfterClick()).
		Row().
		Button(fmt.Sprintf("Автоопределение языка: %v", getAutoDetectChar(ctx, update.Message.From.ID)), []byte("lang-autodetect"), onInlineKeyboardSelect)

	return kb
}

func onInlineKeyboardSelect(ctx context.Context, b *bot.Bot, mes models.MaybeInaccessibleMessage, data []byte) {
	if string(data) == "lang-autodetect" {
		utils.ChangeAutoDetect(ctx, mes.Message.Chat.ID)

		newText := fmt.Sprintf("Автоопределение языка: %v", getAutoDetectChar(ctx, mes.Message.Chat.ID))
		newKb := inline.New(b, inline.NoDeleteAfterClick()).
			Row().
			Button(newText, []byte("lang-autodetect"), onInlineKeyboardSelect)

		_, err := b.EditMessageReplyMarkup(ctx, &bot.EditMessageReplyMarkupParams{
			ChatID:      mes.Message.Chat.ID,
			MessageID:   mes.Message.ID,
			ReplyMarkup: newKb,
		})
		if err != nil {
			log.Println("Error when updating keyboard:", err)
		}
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
