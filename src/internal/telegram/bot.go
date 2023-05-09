package telegram

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"src/internal/services"
)

type State int

const (
	DEFAULT State = iota
	SET
	GET
	DELETE
)

type Bot struct {
	api        *tgbotapi.BotAPI
	service    services.PasswordRecordService
	userStates map[string]State
	messages   []tgbotapi.Message
}

func NewBot(api *tgbotapi.BotAPI, service services.PasswordRecordService) *Bot {
	return &Bot{
		api:        api,
		service:    service,
		userStates: map[string]State{},
		messages:   []tgbotapi.Message{},
	}
}

func (b *Bot) Start() error {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	updates := b.api.GetUpdatesChan(u)
	for update := range updates {
		if update.Message == nil {
			continue
		}

		res := b.HandleAny(update.Message)
		res.ParseMode = "MarkdownV2"
		sent, err := b.api.Send(res)
		if err != nil {
			log.Println("sent failed")
		}
		b.messages = append(b.messages, sent)

	}
	return nil
}
