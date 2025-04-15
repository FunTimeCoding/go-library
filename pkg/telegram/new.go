package telegram

import "github.com/funtimecoding/go-library/pkg/telegram/client"

func New(token string) *Client {
	return &Client{client: client.New(token)}
}
