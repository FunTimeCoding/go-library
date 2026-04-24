package telegram

import "github.com/go-telegram-bot-api/telegram-bot-api/v5"

func (c *Client) SendMessage(
	chat int64,
	s string,
) {
	c.Send(tgbotapi.NewMessage(chat, s))
}
