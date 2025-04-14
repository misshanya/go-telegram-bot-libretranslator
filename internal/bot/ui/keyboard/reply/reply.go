package reply

import (
	"context"
	"github.com/misshanya/go-telegram-bot-libretranslator/internal/bot/ui/keyboard/inline"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
	"github.com/go-telegram/ui/keyboard/reply"
)

type ReplyKeyboard struct {
	Inline *inline.InlineKeyboard
	Reply  *reply.ReplyKeyboard
}

func NewReplyKeyboard(inline *inline.InlineKeyboard) *ReplyKeyboard {
	return &ReplyKeyboard{Inline: inline}
}

func (k *ReplyKeyboard) InitReplyKeyboard(b *bot.Bot) *reply.ReplyKeyboard {
	replyKeyboard := reply.New(
		reply.WithPrefix("reply_keyboard"),
		reply.IsSelective(),
		reply.ResizableKeyboard(),
		reply.IsOneTimeKeyboard(),
	).
		Button("Меню", b, bot.MatchTypeExact, k.onReplyKeyboardMenu)

	k.Reply = replyKeyboard

	return replyKeyboard
}

func (k *ReplyKeyboard) onReplyKeyboardMenu(ctx context.Context, b *bot.Bot, update *models.Update) {
	b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID:      update.Message.Chat.ID,
		Text:        "Меню",
		ReplyMarkup: k.Inline.InitInlineKeyboard(ctx, b, update),
	})
}
