package telegram

import (
	"github.com/funtimecoding/go-library/pkg/bolt"
	"github.com/funtimecoding/go-library/pkg/chat/telegram/client"
)

func New(
	token string,
	database string,
) *Client {
	// https://github.com/go-telegram-bot-api/telegram-bot-api
	result := &Client{client: client.New(token)}

	if database != "" {
		result.setDatabase(bolt.New(database))
	}

	return result
}
