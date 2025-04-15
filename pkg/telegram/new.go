package telegram

import "github.com/funtimecoding/go-library/pkg/telegram/client"

func New(token string) *Client {
	// https://github.com/go-telegram-bot-api/telegram-bot-api
	return &Client{client: client.New(token)}
}
