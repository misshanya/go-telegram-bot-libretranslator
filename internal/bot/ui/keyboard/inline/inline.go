package inline

import (
	"context"
	"fmt"
	"github.com/misshanya/go-telegram-bot-libretranslator/internal/service"
	"log"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
	"github.com/go-telegram/ui/keyboard/inline"
)

type InlineKeyboard struct {
	s service.Service
}

func NewInlineKeyboard(service service.Service) *InlineKeyboard {
	return &InlineKeyboard{s: service}
}

func (k *InlineKeyboard) InitInlineKeyboard(ctx context.Context, b *bot.Bot, update *models.Update) *inline.Keyboard {
	kb, err := k.createKeyboard(ctx, b, update.Message.From.ID)
	if err != nil {
		log.Println(err)
		return nil
	}
	return kb
}

func (k *InlineKeyboard) onInlineKeyboardSelect(ctx context.Context, b *bot.Bot, mes models.MaybeInaccessibleMessage, data []byte) {
	switch string(data) {
	case "lang-autodetect":
		k.s.ChangeAutoDetect(ctx, mes.Message.Chat.ID)
		k.updateKeyboard(ctx, b, &mes)

	case "source-lang":
		err := k.s.SwitchSourceLang(ctx, mes.Message.Chat.ID)
		if err != nil {
			b.SendMessage(ctx, &bot.SendMessageParams{
				ChatID: mes.Message.Chat.ID,
				Text:   "Не удалось изменить язык оригинала",
			})
		}

		k.updateKeyboard(ctx, b, &mes)

	case "target-lang":
		err := k.s.SwitchTargetLang(ctx, mes.Message.Chat.ID)
		if err != nil {
			b.SendMessage(ctx, &bot.SendMessageParams{
				ChatID: mes.Message.Chat.ID,
				Text:   "Не удалось изменить язык оригинала",
			})
		}

		k.updateKeyboard(ctx, b, &mes)
	}
}

func (k *InlineKeyboard) createKeyboard(ctx context.Context, b *bot.Bot, uid int64) (*inline.Keyboard, error) {
	autoDetectChar, autoDetect := k.getAutoDetectChar(ctx, uid)
	autoDetectText := fmt.Sprintf("Автоопределение языка: %v", autoDetectChar)
	sourceLang, err := k.s.GetSourceLang(ctx, uid)
	if err != nil {
		return nil, err
	}
	targetLang, err := k.s.GetTargetLang(ctx, uid)
	if err != nil {
		return nil, err
	}
	sourceLangText := fmt.Sprintf("Переводить с: %v", getFullLangName(sourceLang))
	targetLangText := fmt.Sprintf("Переводить на: %v", getFullLangName(targetLang))

	// Basic keyboard
	kb := inline.New(b, inline.NoDeleteAfterClick()).
		Row().
		Button(autoDetectText, []byte("lang-autodetect"), k.onInlineKeyboardSelect)

	// Add language options if autodetect is false
	if !autoDetect {
		kb = kb.Row().
			Button(sourceLangText, []byte("source-lang"), k.onInlineKeyboardSelect).
			Button(targetLangText, []byte("target-lang"), k.onInlineKeyboardSelect)
	}

	return kb, nil
}

func (k *InlineKeyboard) updateKeyboard(ctx context.Context, b *bot.Bot, mes *models.MaybeInaccessibleMessage) error {
	newKb, err := k.createKeyboard(ctx, b, mes.Message.Chat.ID)
	if err != nil {
		return err
	}

	_, err = b.EditMessageReplyMarkup(ctx, &bot.EditMessageReplyMarkupParams{
		ChatID:      mes.Message.Chat.ID,
		MessageID:   mes.Message.ID,
		ReplyMarkup: newKb,
	})
	if err != nil {
		log.Println("Error when updating keyboard:", err)
	}
	return nil
}

func (k *InlineKeyboard) getAutoDetectChar(ctx context.Context, uid int64) (string, bool) {
	var autoDetectChar string
	autoDetect := k.s.IsAutoDetect(ctx, uid)
	if autoDetect {
		autoDetectChar = "✅"
	} else {
		autoDetectChar = "❎"
	}
	return autoDetectChar, autoDetect
}

func getFullLangName(lang string) string {
	switch lang {
	case "ru":
		return "Русский"
	case "en":
		return "Английский"
	}
	return ""
}
