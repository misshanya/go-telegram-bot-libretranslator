package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/misshanya/go-telegram-bot-libretranslator/internal/config"
)

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

func IsAutoDetect(uid int) bool {
	// todo: realize functional to check user's config
	return true
}
