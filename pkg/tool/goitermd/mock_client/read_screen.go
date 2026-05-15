package mock_client

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/iterm/session"
)

func (c *Client) ReadScreen(identifier string) (session.Screen, error) {
	s, okay := c.screens[identifier]

	if !okay {
		return session.Screen{}, fmt.Errorf(
			"session %s not found",
			identifier,
		)
	}

	return s, nil
}
