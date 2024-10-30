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
	"github.com/misshanya/go-telegram-bot-libretranslator/internal/db/users"
)

type LanguageDetection struct {
	Confidence float64 `json:"confidence"`
	Language   string  `json:"language"`
}

// postRequest sends a POST request to the specified URL with the given data
// serialized as JSON and returns the response body as a byte slice.
// It returns an error if any step of the process fails.
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

// Translate sends a request to the LibreTranslate API to translate the given text
// from the source language to the target language and returns the translated text.
func Translate(text string, sourceLang string, targetLang string) (string, error) {
	postBody := map[string]string{
		"q":      text,
		"source": sourceLang,
		"target": targetLang,
	}
	url := fmt.Sprintf("%v/translate", config.GetConfig().LibreTranslateUrl)
	body, err := postRequest(url, postBody)
	if err != nil {
		return "", err
	}

	var result map[string]string
	err = json.Unmarshal(body, &result)
	if err != nil {
		return "", err
	}

	translatedText := result["translatedText"]
	return translatedText, nil
}

// DetectLanguage sends a request to the LibreTranslate API to detect the language of the given text
// and returns the detected language as a string.
func DetectLanguage(text string) (string, error) {
	postBody := map[string]string{
		"q": text,
	}
	url := fmt.Sprintf("%v/detect", config.GetConfig().LibreTranslateUrl)
	body, err := postRequest(url, postBody)
	if err != nil {
		return "", err
	}

	var result []LanguageDetection
	err = json.Unmarshal(body, &result)
	if err != nil {
		return "", err
	}

	detectedLang := result[0].Language
	return detectedLang, nil
}

// IsAutoDetect checks if the user's language autodetect feature is enabled.
// It returns true if enabled, false otherwise.
func IsAutoDetect(ctx context.Context, uid int64) bool {
	queries := config.GetDB()
	user, err := queries.GetUser(ctx, uid)
	if err != nil {
		return false
	}
	return user.LangAutodetect
}

// ChangeAutoDetect toggles the language autodetect feature for the user with the given UID.
// It logs an error if the operation fails.
func ChangeAutoDetect(ctx context.Context, uid int64) {
	queries := config.GetDB()
	_, err := queries.ChangeLangAutodetect(ctx, uid)
	if err != nil {
		log.Println(err)
	}
}

// RegisterUser attempts to register a new user with the given UID.
// It returns true if the user was successfully registered, or false if the user already exists.
func RegisterUser(ctx context.Context, uid int64) bool {
	queries := config.GetDB()
	_, err := queries.GetUser(ctx, uid)
	if err == nil {
		return false
	}
	_, err = queries.CreateUser(ctx, uid)
	return err != nil
}

func GetSourceLang(ctx context.Context, uid int64) (string, error) {
	queries := config.GetDB()
	sourceLang, err := queries.GetSourceLang(ctx, uid)
	if err != nil {
		log.Println(err)
		return "", err
	}
	return sourceLang, nil
}

func GetTargetLang(ctx context.Context, uid int64) (string, error) {
	queries := config.GetDB()
	sourceLang, err := queries.GetTargetLang(ctx, uid)
	if err != nil {
		log.Println(err)
		return "", err
	}
	return sourceLang, nil
}

func SetSourceLang(ctx context.Context, uid int64, newSourceLang string) error {
	queries := config.GetDB()
	err := queries.SetSourceLang(ctx, users.SetSourceLangParams{
		SourceLang: newSourceLang,
		TgID:       uid,
	},
	)
	return err
}

func SetTargetLang(ctx context.Context, uid int64, newTargetLang string) error {
	queries := config.GetDB()
	err := queries.SetTargetLang(ctx, users.SetTargetLangParams{
		TargetLang: newTargetLang,
		TgID:       uid,
	},
	)
	return err
}

func GetOppositeLang(lang string) string {
	switch lang {
	case "ru":
		return "en"
	case "en":
		return "ru"
	}
	return ""
}
