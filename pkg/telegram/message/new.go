package message

import "github.com/go-telegram-bot-api/telegram-bot-api/v5"

func New(m *tgbotapi.Message) *Message {
	return &Message{From: m.From.UserName, Text: m.Text, Raw: m}
}
