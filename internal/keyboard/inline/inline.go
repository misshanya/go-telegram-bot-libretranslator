package inline

import (
	"context"
	"fmt"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
	"github.com/go-telegram/ui/keyboard/inline"
	"github.com/misshanya/go-telegram-bot-libretranslator/internal/utils"
)

func InitInlineKeyboard(b *bot.Bot) *inline.Keyboard {
	kb := inline.New(b, inline.NoDeleteAfterClick()).
		Row().
		Button(fmt.Sprintf("Автоопределение языка: %v", utils.IsAutoDetect(1)), []byte("lang-autodetect"), onInlineKeyboardSelect)

	return kb
}

func onInlineKeyboardSelect(ctx context.Context, b *bot.Bot, mes models.MaybeInaccessibleMessage, data []byte) {
	if string(data) == "lang-autodetect" {
		b.SendMessage(ctx, &bot.SendMessageParams{
			ChatID: mes.Message.Chat.ID,
			Text:   "Нажато " + string(data),
		})
	}
}
