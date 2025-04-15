package telegram

import "github.com/funtimecoding/go-library/pkg/telegram/message"

func (c *Client) Messages() []*message.Message {
	var result []*message.Message

	for _, u := range c.Updates() {
		if u.Message != nil {
			m := message.New(u.Message)
			m.Update = &u
			result = append(result, m)
		}
	}

	return result
}
