package main

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strings"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
	"github.com/joho/godotenv"
)

var LibreTranslateUrl string

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	token := os.Getenv("BOT_TOKEN")
	LibreTranslateUrl = os.Getenv("LIBRETRANSLATE_URL")

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	b, err := bot.New(token)
	if err != nil {
		panic(err)
	}

	b.RegisterHandler(bot.HandlerTypeMessageText, "/start", bot.MatchTypeExact, startHandler)
	b.RegisterHandler(bot.HandlerTypeMessageText, "/translate", bot.MatchTypePrefix, translateHandler)

	b.Start(ctx)
}

func startHandler(ctx context.Context, b *bot.Bot, update *models.Update) {
	b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID: update.Message.Chat.ID,
		Text:   "Привет! Я бот-переводчик. Написан на Go. Для перевода используется API LibreTranslate",
	})
}

func translateHandler(ctx context.Context, b *bot.Bot, update *models.Update) {

	textToTranslate := strings.TrimSpace(strings.TrimPrefix(update.Message.Text, "/translate"))

	translatedText := translate(textToTranslate, "ru", "en")

	b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID: update.Message.Chat.ID,
		Text:   translatedText,
	})
}

func translate(text string, langFrom string, langTo string) string {
	postBody, _ := json.Marshal(map[string]string{
		"q":      text,
		"source": langFrom,
		"target": langTo,
	})
	responseBody := bytes.NewBuffer(postBody)
	resp, err := http.Post(LibreTranslateUrl+"/translate", "application/json", responseBody)
	if err != nil {
		log.Fatalf("An Error Occured %v", err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	var result map[string]string
	err = json.Unmarshal(body, &result)
	if err != nil {
		log.Fatalln(err)
	}

	translatedText := result["translatedText"]
	return translatedText
}
