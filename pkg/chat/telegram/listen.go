package telegram

import "github.com/go-telegram-bot-api/telegram-bot-api/v5"

func (c *Client) Listen() tgbotapi.UpdatesChannel {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	return c.client.GetUpdatesChan(u)
}
