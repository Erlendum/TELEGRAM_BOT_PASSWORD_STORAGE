package flags

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type TelegramBotFlags struct {
	Token string `mapstructure:"token"`
}

func (f *TelegramBotFlags) Init() (*tgbotapi.BotAPI, error) {
	botAPI, err := tgbotapi.NewBotAPI(f.Token)
	if err != nil {
		return nil, err
	}
	return botAPI, nil
}
