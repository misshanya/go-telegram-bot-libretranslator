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

func Translate(text string, langFrom string, langTo string) string {
	postBody, _ := json.Marshal(map[string]string{
		"q":      text,
		"source": langFrom,
		"target": langTo,
	})
	requestBody := bytes.NewBuffer(postBody)
	url := fmt.Sprintf("%v/translate", config.GetConfig().LibreTranslateUrl)
	resp, err := http.Post(url, "application/json", requestBody)
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

func IsAutoDetect(uid int) bool {
	// todo: realize functional to check user's config
	return true
}
