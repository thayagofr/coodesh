package utils

import (
	"log"
	"net/http"
	"net/url"
	"os"
)

func NotifyTelegram(text string) {
	API := os.Getenv("TELEGRAM_URL")
	CHATID := os.Getenv("TELEGRAM_CHAT_ID")
	_ , err := http.PostForm(
		API,
		url.Values{
			"chat_id": {CHATID},
			"text":    {text},
		})
	if err != nil {
		log.Printf("Erro ao notificar problema de sincronizacao")
	}
}
