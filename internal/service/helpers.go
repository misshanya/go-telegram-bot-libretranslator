package service

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/misshanya/go-telegram-bot-libretranslator/internal/config"
)

type languageDetection struct {
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

// translate sends a request to the LibreTranslate API to translate the given text
// from the source language to the target language and returns the translated text.
func translate(text string, sourceLang string, targetLang string) (string, error) {
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

// detectLanguage sends a request to the LibreTranslate API to detect the language of the given text
// and returns the detected language as a string.
func detectLanguage(text string) (string, error) {
	postBody := map[string]string{
		"q": text,
	}
	url := fmt.Sprintf("%v/detect", config.GetConfig().LibreTranslateUrl)
	body, err := postRequest(url, postBody)
	if err != nil {
		return "", err
	}

	var result []languageDetection
	err = json.Unmarshal(body, &result)
	if err != nil {
		return "", err
	}

	detectedLang := result[0].Language
	return detectedLang, nil
}

// getOppositeLang returns the opposite language for the given language
func getOppositeLang(lang string) string {
	switch lang {
	case "ru":
		return "en"
	case "en":
		return "ru"
	}
	return ""
}
