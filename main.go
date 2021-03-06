package main

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

const TOKEN = "1292527443:AAEbZeT9LuXeXQcFtloDH35xmDWtFiqV1Vw"
const MISHA_ID = 426117512

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

		log.Printf("%s %s [%d] > %s", update.Message.From.FirstName, update.Message.From.LastName, update.Message.From.ID, update.Message.Text)

		if update.Message.From.ID == MISHA_ID {
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Миша?")
			msg.ReplyToMessageID = update.Message.MessageID

			bot.Send(msg)
		}
	}
}
