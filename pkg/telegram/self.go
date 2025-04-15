package telegram

import "github.com/go-telegram-bot-api/telegram-bot-api/v5"

func (c *Client) Self() tgbotapi.User {
	return c.client.Self
}
