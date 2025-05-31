package main

import (
	"encoding/json"
	"log"
	"net/http"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5 v5.5.1 h1:wG8n/XJQ07TmjbITcGiUaOtXxdrINDz1b0J1w0SzqDc="
)

func main() {
	http.Handle("/", http.FileServer(http.Dir("./frontend")))

	log.Println("Server running on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}

func handleUpdate(update tgbotapi.Update) {
	if update.Message != nil {
		// The library might not expose WebAppData, so parse raw JSON
		rawJSON := update.Message.Raw // or wherever raw JSON is stored

		var rawMap map[string]interface{}
		err := json.Unmarshal(rawJSON, &rawMap)
		if err != nil {
			log.Println("Failed to parse update JSON:", err)
			return
		}

		if webAppData, ok := rawMap["web_app_data"]; ok {
			log.Println("Received web_app_data:", webAppData)
			// handle web_app_data here
		}
	}
}
