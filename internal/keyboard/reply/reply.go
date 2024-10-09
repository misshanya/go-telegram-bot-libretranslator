package reply

import (
	"context"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
	"github.com/go-telegram/ui/keyboard/reply"
	"github.com/misshanya/go-telegram-bot-libretranslator/internal/keyboard/inline"
)

func InitReplyKeyboard(b *bot.Bot) *reply.ReplyKeyboard {
	demoReplyKeyboard := reply.New(
		reply.WithPrefix("reply_keyboard"),
		reply.IsSelective(),
		reply.ResizableKeyboard(),
		reply.IsOneTimeKeyboard(),
	).
		Button("Меню", b, bot.MatchTypeExact, onReplyKeyboardMenu)

	return demoReplyKeyboard
}

func onReplyKeyboardMenu(ctx context.Context, b *bot.Bot, update *models.Update) {
	b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID:      update.Message.Chat.ID,
		Text:        "Меню",
		ReplyMarkup: inline.InitInlineKeyboard(ctx, b, update),
	})
}
