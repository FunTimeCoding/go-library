package message

import "github.com/go-telegram-bot-api/telegram-bot-api/v5"

type Message struct {
	From   string
	Text   string
	Raw    *tgbotapi.Message
	Update *tgbotapi.Update
}
