package mock_client

import "github.com/funtimecoding/go-library/pkg/iterm/screen"

func New() *Client {
	return &Client{
		screens: map[string]*screen.Screen{},
		nextID:  1,
	}
}
