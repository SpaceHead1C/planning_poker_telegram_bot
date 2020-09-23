package main

import (
	"fmt"
	"net/http"
	"time"

	"./betypes"
	"./logger"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

var (
	NewBot, BotErr = tgbotapi.NewBotAPI(betypes.BOT_TOKEN)
)

func setWebhook(bot *tgbotapi.BotAPI) {
	webHookInfo := tgbotapi.NewWebhookWithCert(fmt.Sprintf("https://%s:%s/%s", betypes.BOT_ADDRESS, betypes.BOT_PORT, betypes.BOT_TOKEN), betypes.CERT_PATH)
	_, err := bot.SetWebhook(webHookInfo)
	logger.ForError(err)
}

func main() {
	logger.ForError(BotErr)

	fmt.Println("OK", time.Now().Unix(), time.Now(), time.Now().Weekday())

	setWebhook(NewBot)

	updates := NewBot.ListenForWebhook("/" + betypes.BOT_TOKEN)

	go http.ListenAndServeTLS(fmt.Sprintf("%s:%s", betypes.BOT_ADDRESS, betypes.BOT_PORT), betypes.CERT_PATH, betypes.KEY_PATH, nil)

	for update := range updates {
		if update.Message == nil {
			continue
		}

		m := update.Message
		fmt.Println(m.From.UserName, m.From.ID, m.Chat.ID, m.MessageID, m.Text)
	}
}
