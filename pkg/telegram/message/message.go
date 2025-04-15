package message

import (
	"github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"time"
)

type Message struct {
	From   string
	Text   string
	Create *time.Time

	Update *tgbotapi.Update

	Raw *tgbotapi.Message
}
