package main

import (
	"log"
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func main() {
	botToken := os.Getenv("TELEGRAM_BOT_TOKEN")
	if botToken == "" {
		log.Fatal("TELEGRAM_BOT_TOKEN is not set")
	}

	bot, err := tgbotapi.NewBotAPI(botToken)
	if err != nil {
		log.Panic(err)
	}

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message != nil {
			chatID := update.Message.Chat.ID

			if update.Message.WebAppData != nil {
				// Received data from Mini App
				log.Printf("Received WebApp data: %s", update.Message.WebAppData.Data)

				msg := tgbotapi.NewMessage(chatID, "Thanks for the data: "+update.Message.WebAppData.Data)
				bot.Send(msg)
				continue
			}

			// Send message with InlineKeyboardButton that opens Mini App
			webAppInfo := &tgbotapi.WebAppInfo{URL: "http://localhost:8080"}

			btn := tgbotapi.NewInlineKeyboardButtonWebApp("Open Mini App", webAppInfo)
			keyboard := tgbotapi.NewInlineKeyboardMarkup(
				tgbotapi.NewInlineKeyboardRow(btn),
			)

			msg := tgbotapi.NewMessage(chatID, "Please open the Mini App")
			msg.ReplyMarkup = keyboard

			bot.Send(msg)
		}
	}
}
