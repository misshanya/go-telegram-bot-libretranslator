package utils

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/misshanya/go-telegram-bot-libretranslator/internal/config"
)

type LanguageDetection struct {
	Confidence float64 `json:"confidence"`
	Language   string  `json:"language"`
}

func postRequest(url string, data interface{}) ([]byte, error) {
	jsonData, err := json.Marshal(data)
	if err != nil {
		return nil, fmt.Errorf("error marshaling data: %w", err)
	}

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, fmt.Errorf("error requesting: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response: %w", err)
	}

	return body, nil
}

func Translate(text string, langFrom string, langTo string) string {
	postBody := map[string]string{
		"q":      text,
		"source": langFrom,
		"target": langTo,
	}
	url := fmt.Sprintf("%v/translate", config.GetConfig().LibreTranslateUrl)
	body, err := postRequest(url, postBody)
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

func DetectLanguage(text string) string {
	postBody := map[string]string{
		"q": text,
	}
	url := fmt.Sprintf("%v/detect", config.GetConfig().LibreTranslateUrl)
	body, err := postRequest(url, postBody)
	if err != nil {
		log.Println("Error when requesting detect:", err)
	}

	var result []LanguageDetection
	err = json.Unmarshal(body, &result)
	if err != nil {
		log.Println("Error when unmarshaling detectlang response:", err)
	}

	detectedLang := result[0].Language
	return detectedLang
}

func IsAutoDetect(ctx context.Context, uid int64) bool {
	queries := config.GetDB()
	user, err := queries.GetUser(ctx, uid)
	if err != nil {
		return false
	}
	return user.LangAutodetect
}

func ChangeAutoDetect(ctx context.Context, uid int64) {
	queries := config.GetDB()
	_, err := queries.ChangeLangAutodetect(ctx, uid)
	if err != nil {
		log.Println(err)
	}
}

func RegisterUser(ctx context.Context, uid int64) bool {
	queries := config.GetDB()
	_, err := queries.GetUser(ctx, uid)
	if err == nil {
		return false
	}
	_, err = queries.CreateUser(ctx, uid)
	return err != nil
}
