package message

import "github.com/go-telegram-bot-api/telegram-bot-api/v5"

func NewSlice(v []*tgbotapi.Message) []*Message {
	var result []*Message

	for _, m := range v {
		result = append(result, New(m))
	}

	return result
}
