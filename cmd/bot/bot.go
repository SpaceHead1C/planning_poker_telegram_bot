package bot

import (
	"fmt"
	"net/http"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

type Bot struct {
	BotAPI   *tgbotapi.BotAPI
	updates  tgbotapi.UpdatesChannel
	delivery chan *tgbotapi.MessageConfig
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

	b := &Bot{
		BotAPI:   botAPI,
		updates:  updates,
		delivery: make(chan *tgbotapi.MessageConfig),
	}

	go b.send()
	go b.listen()

	return b
}

func (b *Bot) String() string {
	return "Planning Poker"
}

func (b *Bot) listen() {
	for update := range b.updates {
		if update.Message == nil {
			continue
		}

		msg := update.Message

		b.processMessage(msg)

		msgLog := messageForLog(msg)
		fmt.Println(msgLog)
	}
}

func (b *Bot) send() {
	for msg := range b.delivery {
		if _, err := b.BotAPI.Send(*msg); err != nil {
			fmt.Println(err)
		}
	}
}

func (b *Bot) processMessage(m *tgbotapi.Message) {
	if m.IsCommand() {
		b.processComand(m)
	}
}

func (b *Bot) SendStartMenu(id int64, text string) {
	msg := tgbotapi.NewMessage(id, "Комнаты")
	b.delivery <- &msg
}

func (b *Bot) processComand(m *tgbotapi.Message) {
	switch m.Command() {
	case "start":
		fallthrough
	case "menu":
		b.SendStartMenu(m.Chat.ID, "Комнаты")
	default:
		msg := tgbotapi.NewMessage(m.Chat.ID, "Неизвестная команда")
		b.delivery <- &msg
	}
}

func messageForLog(m *tgbotapi.Message) string {
	return fmt.Sprintf(`Сообщение:
	ID: %d
	Text: %s
	Command: %s
	IsCommand: %v
	Чат:
		ID: %d
		Title: %s
		Description: %s
		Username: %s
		Firstname: %s
		Lastname: %s
		InviteLink: %s
		Type: %s
	От кого:
		ID: %d
		IsBot: %v
		Username: %s
		Firstname: %s
		Lastname: %s
		LanguageCode: %s
	Дата: %d
	Поля: %+v
`,
		m.MessageID,
		m.Text,
		m.Command(),
		m.IsCommand(),
		m.Chat.ID,
		m.Chat.Title,
		m.Chat.Description,
		m.Chat.UserName,
		m.Chat.FirstName,
		m.Chat.LastName,
		m.Chat.InviteLink,
		m.Chat.Type,
		m.From.ID,
		m.From.IsBot,
		m.From.UserName,
		m.From.FirstName,
		m.From.LastName,
		m.From.LanguageCode,
		m.Date,
		m.Entities,
	)
}
