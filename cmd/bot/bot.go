package bot

import (
	"fmt"
	"net/http"
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

type Bot struct {
	BotAPI  *tgbotapi.BotAPI
	Updates tgbotapi.UpdatesChannel
}

func NewBot(c *Config) *Bot {
	botAPI, err := tgbotapi.NewBotAPI(c.BotToken)
	if err != nil {
		panic(err.Error())
	}

	webHookInfo := tgbotapi.NewWebhookWithCert(fmt.Sprintf("https://%s:%s/%s", c.BotAddress, c.BotPort, c.BotToken), c.CertPath)
	if _, err := botAPI.SetWebhook(webHookInfo); err != nil {
		panic(err.Error())
	}

	updates := botAPI.ListenForWebhook("/" + c.BotToken)

	go http.ListenAndServeTLS(fmt.Sprintf("%s:%s", c.BotAddress, c.BotPort), c.CertPath, c.KeyPath, nil)

	return &Bot{
		BotAPI:  botAPI,
		Updates: updates,
	}
}

func (b *Bot) Listen(c chan<- string) {
	for update := range b.Updates {
		if update.Message == nil {
			fmt.Println("empty message")
			continue
		}

		m := update.Message
		msg := fmt.Sprintln(m.From.UserName, m.From.ID, m.Chat.ID, m.MessageID, m.Text)

		b.BotAPI.Send(tgbotapi.NewMessage(m.Chat.ID, msg))

		c <- msg
	}
}

func (b *Bot) UserID(m *tgbotapi.Message) string {
	return "telegram|" + strconv.Itoa(m.From.ID)
}
