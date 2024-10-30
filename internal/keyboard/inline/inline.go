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
	kb, err := createKeyboard(ctx, b, update.Message.From.ID)
	if err != nil {
		log.Println(err)
		return nil
	}
	return kb
}

func onInlineKeyboardSelect(ctx context.Context, b *bot.Bot, mes models.MaybeInaccessibleMessage, data []byte) {
	switch string(data) {
	case "lang-autodetect":
		utils.ChangeAutoDetect(ctx, mes.Message.Chat.ID)
		updateKeyboard(ctx, b, &mes)

	case "source-lang":
		currentSourceLang, err := utils.GetSourceLang(ctx, mes.Message.Chat.ID)
		if err != nil {
			log.Println(err)
		}
		newSourceLang := utils.GetOppositeLang(currentSourceLang)
		err = utils.SetSourceLang(ctx, mes.Message.Chat.ID, newSourceLang)
		if err != nil {
			b.SendMessage(ctx, &bot.SendMessageParams{
				ChatID: mes.Message.Chat.ID,
				Text:   "Не удалось изменить язык оригинала",
			})
		}
		updateKeyboard(ctx, b, &mes)

	case "target-lang":
		currentTargetLang, err := utils.GetTargetLang(ctx, mes.Message.Chat.ID)
		if err != nil {
			log.Println(err)
		}
		newTargetLang := utils.GetOppositeLang(currentTargetLang)
		err = utils.SetTargetLang(ctx, mes.Message.Chat.ID, newTargetLang)
		if err != nil {
			b.SendMessage(ctx, &bot.SendMessageParams{
				ChatID: mes.Message.Chat.ID,
				Text:   "Не удалось изменить язык перевода",
			})
		}
		updateKeyboard(ctx, b, &mes)
	}
}

func createKeyboard(ctx context.Context, b *bot.Bot, uid int64) (*inline.Keyboard, error) {
	autoDetectChar, autoDetect := getAutoDetectChar(ctx, uid)
	autoDetectText := fmt.Sprintf("Автоопределение языка: %v", autoDetectChar)
	sourceLang, err := utils.GetSourceLang(ctx, uid)
	if err != nil {
		return nil, err
	}
	targetLang, err := utils.GetTargetLang(ctx, uid)
	if err != nil {
		return nil, err
	}
	sourceLangText := fmt.Sprintf("Переводить с: %v", getFullLangName(sourceLang))
	targetLangText := fmt.Sprintf("Переводить на: %v", getFullLangName(targetLang))

	// Basic keyboard
	kb := inline.New(b, inline.NoDeleteAfterClick()).
		Row().
		Button(autoDetectText, []byte("lang-autodetect"), onInlineKeyboardSelect)

	// Add language options if autodetect is false
	if !autoDetect {
		kb = kb.Row().
			Button(sourceLangText, []byte("source-lang"), onInlineKeyboardSelect).
			Button(targetLangText, []byte("target-lang"), onInlineKeyboardSelect)
	}

	return kb, nil
}

func updateKeyboard(ctx context.Context, b *bot.Bot, mes *models.MaybeInaccessibleMessage) error {
	newKb, err := createKeyboard(ctx, b, mes.Message.Chat.ID)
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

func getAutoDetectChar(ctx context.Context, uid int64) (string, bool) {
	var autoDetectChar string
	autoDetect := utils.IsAutoDetect(ctx, uid)
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
