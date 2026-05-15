package mock_client

import "github.com/funtimecoding/go-library/pkg/iterm/session"

func New() *Client {
	return &Client{
		screens: map[string]session.Screen{},
		nextID:  1,
	}
}
