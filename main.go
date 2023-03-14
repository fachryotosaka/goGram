package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func main() {
	token := "6071514173:AAFJIPRZ4w5QnxlQTpIJIbjrnvsmcIw7lxI"
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message != nil { // If we got a message
			log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

			if strings.Contains(update.Message.Text, "ready") {
				message, err := os.ReadFile("template/ready.txt")
				if err != nil {
					log.Fatal(err)
				}

				msg := tgbotapi.NewMessage(update.Message.Chat.ID, string(message))
				msg.ReplyToMessageID = update.Message.MessageID

				bot.Send(msg)
			} else if strings.Contains(update.Message.Text, "price") {
				message, err := os.ReadFile("template/price.txt")
				if err != nil {
					log.Fatal(err)
				}

				msg := tgbotapi.NewMessage(update.Message.Chat.ID, string(message))
				msg.ReplyToMessageID = update.Message.MessageID
				bot.Send(msg)

			} else if strings.Contains(update.Message.Text, "who") {
				message, err := os.ReadFile("template/who.txt")
				if err != nil {
					log.Fatal(err)
				}

				msg := tgbotapi.NewMessage(update.Message.Chat.ID, string(message))
				msg.ReplyToMessageID = update.Message.MessageID
				bot.Send(msg)

			} else if strings.Contains(update.Message.Text, "lexa") {
				message, err := os.ReadFile("template/lexa.txt")
				if err != nil {
					log.Fatal(err)
				}

				msg := tgbotapi.NewMessage(update.Message.Chat.ID, string(message))
				msg.ReplyToMessageID = update.Message.MessageID
				bot.Send(msg)

			} else {
				welcomeMessage := fmt.Sprintf("Hello %s, ask some question...", update.Message.From.UserName)
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, string(welcomeMessage))
				msg.ReplyToMessageID = update.Message.MessageID
				bot.Send(msg)
			}
		}
	}
}
