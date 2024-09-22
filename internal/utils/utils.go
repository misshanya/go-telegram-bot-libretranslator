package utils

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
)

func Translate(text string, langFrom string, langTo string) string {
	postBody, _ := json.Marshal(map[string]string{
		"q":      text,
		"source": langFrom,
		"target": langTo,
	})
	responseBody := bytes.NewBuffer(postBody)
	resp, err := http.Post("http://127.0.0.1:8000/translate", "application/json", responseBody)
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
	return true
}
