package telegram

import (
	"fmt"
	"github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (c *Client) Example() {
	c.client.Debug = true
	fmt.Printf("User: %s\n", c.client.Self.UserName)
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	for p := range c.client.GetUpdatesChan(u) {
		if p.Message == nil {
			continue
		}

		fmt.Printf("%s %s", p.Message.From.UserName, p.Message.Text)
		m := tgbotapi.NewMessage(p.Message.Chat.ID, p.Message.Text)
		m.ReplyToMessageID = p.Message.MessageID
		c.Send(m)
	}
}
