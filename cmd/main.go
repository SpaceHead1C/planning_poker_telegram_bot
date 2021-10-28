package main

import (
	"fmt"
	"net/http"
	"time"

	"telegram-bot-planning-poker/cmd/bot"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func main() {
	conf := bot.NewConfig()
	if err := bot.ParseConfig(conf); err != nil {
		panic(err.Error())
	}
	fmt.Println(conf)

	botAPI, err := tgbotapi.NewBotAPI(conf.BotToken)
	if err != nil {
		panic(err.Error())
	}

	fmt.Println("OK", time.Now().Unix(), time.Now(), time.Now().Weekday())

	webHookInfo := tgbotapi.NewWebhookWithCert(fmt.Sprintf("https://%s:%s/%s", conf.BotAddress, conf.BotPort, conf.BotToken), conf.CertPath)
	if _, err := botAPI.SetWebhook(webHookInfo); err != nil {
		panic(err.Error())
	}

	updates := botAPI.ListenForWebhook("/" + conf.BotToken)

	go http.ListenAndServeTLS(fmt.Sprintf("%s:%s", conf.BotAddress, conf.BotPort), conf.CertPath, conf.KeyPath, nil)

	for update := range updates {
		if update.Message == nil {
			fmt.Println("empty message")
			continue
		}

		m := update.Message
		fmt.Println(m.From.UserName, m.From.ID, m.Chat.ID, m.MessageID, m.Text)
	}
}
