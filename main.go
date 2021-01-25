package main

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

const TOKEN = "1292527443:AAEbZeT9LuXeXQcFtloDH35xmDWtFiqV1Vw"

func main() {
	bot, err := tgbotapi.NewBotAPI(TOKEN)
	if err != nil {
		panic(err)
	}

	log.Printf("Authorized on account %s", bot.Self.UserName)

	config := tgbotapi.NewUpdate(0)
	config.Timeout = 60

	updates, err := bot.GetUpdatesChan(config)

	for update := range updates {
		if update.Message == nil { // ignore any non-Message Updates
			continue
		}

		if update.Message.From.FirstName == "м" && update.Message.From.LastName == "б" {
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Миша?")
			msg.ReplyToMessageID = update.Message.MessageID

			bot.Send(msg)
		}
	}
}
