package telegram

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"src/internal/models"
	"src/internal/pkg"
	"strings"
)

const MAXLEN = 100

func (b *Bot) startCommand(msg *tgbotapi.Message) tgbotapi.MessageConfig {
	res := tgbotapi.NewMessage(msg.Chat.ID, pkg.MsgStart)
	return res
}

func (b *Bot) setCommand(msg *tgbotapi.Message) tgbotapi.MessageConfig {
	res := tgbotapi.NewMessage(msg.Chat.ID, pkg.MsgSet)
	b.userStates[msg.From.UserName] = SET
	return res
}

func (b *Bot) getCommand(msg *tgbotapi.Message) tgbotapi.MessageConfig {
	res := tgbotapi.NewMessage(msg.Chat.ID, pkg.MsgGet)
	b.userStates[msg.From.UserName] = GET
	return res
}

func (b *Bot) deleteCommand(msg *tgbotapi.Message) tgbotapi.MessageConfig {
	res := tgbotapi.NewMessage(msg.Chat.ID, pkg.MsgDel)

	b.userStates[msg.From.UserName] = DELETE
	return res
}

func (b *Bot) unknownCommand(message *tgbotapi.Message) tgbotapi.MessageConfig {
	res := tgbotapi.NewMessage(message.Chat.ID, pkg.MsgUnknownCommand)
	return res
}

func (b *Bot) HandleAny(msg *tgbotapi.Message) tgbotapi.MessageConfig {
	b.messages = append(b.messages, *msg)
	usr := msg.From.UserName
	if state, ok := b.userStates[usr]; ok {
		return b.handleState(state, msg)
	}
	if msg.IsCommand() {
		return b.handleCommand(msg)
	}
	return b.unknownCommand(msg)
}

func (b *Bot) handleCommand(msg *tgbotapi.Message) tgbotapi.MessageConfig {
	switch msg.Command() {
	case pkg.CommandStart:
		return b.startCommand(msg)
	case pkg.CommandSet:
		return b.setCommand(msg)
	case pkg.CommandGet:
		return b.getCommand(msg)
	case pkg.CommandDelete:
		return b.deleteCommand(msg)
	default:
		return b.unknownCommand(msg)
	}
}

func (b *Bot) handleState(state State, msg *tgbotapi.Message) tgbotapi.MessageConfig {
	delete(b.userStates, msg.From.UserName)
	switch state {
	case SET:
		return b.set(msg)
	case GET:
		return b.get(msg)
	case DELETE:
		return b.delete(msg)
	default:
		return tgbotapi.NewMessage(msg.Chat.ID, pkg.MsgErrorUnknown)
	}
}

func (b *Bot) set(msg *tgbotapi.Message) tgbotapi.MessageConfig {
	arr := strings.Fields(msg.Text)
	if len(arr) < 2 {
		res := tgbotapi.NewMessage(msg.Chat.ID, pkg.MsgErrorSetInvalidArgs)
		return res
	}

	service := strings.Join(arr[:len(arr)-1], " ")
	password := arr[len(arr)-1]
	if len(service) > MAXLEN || len(password) > MAXLEN {
		res := tgbotapi.NewMessage(msg.Chat.ID, pkg.MsgErrorSetTooLong)
		return res
	}
	err := b.service.Set(&models.PasswordRecord{
		Service:  service,
		Password: password,
	})
	if err != nil {
		res := tgbotapi.NewMessage(msg.Chat.ID, pkg.MsgRepositoryError)
		return res
	}
	text := fmt.Sprintf(pkg.MsgSuccessSet, esc(password), esc(service))
	res := tgbotapi.NewMessage(msg.Chat.ID, text)
	return res
}

func (b *Bot) get(msg *tgbotapi.Message) tgbotapi.MessageConfig {
	service := strings.TrimSpace(msg.Text)
	val, err := b.service.Get(msg.From.UserName, service)
	switch {
	case err != nil:
		res := tgbotapi.NewMessage(msg.Chat.ID, esc(pkg.MsgRepositoryError))
		return res
	case val == nil:
		text := fmt.Sprintf(esc(pkg.MsgErrorGet), esc(msg.Text))
		res := tgbotapi.NewMessage(msg.Chat.ID, text)
		return res
	}
	res := tgbotapi.NewMessage(msg.Chat.ID, esc(val.Password))
	res.ReplyToMessageID = msg.MessageID
	return res
}

func (b *Bot) delete(msg *tgbotapi.Message) tgbotapi.MessageConfig {
	service := strings.TrimSpace(msg.Text)
	err := b.service.Delete(msg.From.UserName, service)
	if err != nil {
		res := tgbotapi.NewMessage(msg.Chat.ID, pkg.MsgErrorDel)
		return res
	}

	text := fmt.Sprintf(pkg.MsgSuccessDel, esc(msg.Text))
	res := tgbotapi.NewMessage(msg.Chat.ID, text)
	return res
}

func esc(s string) string {
	return tgbotapi.EscapeText(tgbotapi.ModeMarkdownV2, s)
}
