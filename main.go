package main

import (
	"log"
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	tgbot, err := tgbotapi.NewBotAPI(os.Getenv("BOT_API"))
	if err != nil {
		log.Fatal("error: ", err)
	}

	up := tgbotapi.NewUpdate(0)
	up.Timeout = 60

	msges, err := tgbot.GetUpdates(up)
}
