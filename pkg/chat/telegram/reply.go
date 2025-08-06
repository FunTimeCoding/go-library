package telegram

import "github.com/go-telegram-bot-api/telegram-bot-api/v5"

func (c *Client) Reply(
	p *tgbotapi.Update,
	s string,
) {
	m := tgbotapi.NewMessage(p.Message.Chat.ID, s)
	m.ReplyToMessageID = p.Message.MessageID
	c.Send(m)
}
