package mock_client

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/iterm/session"
)

func (c *Client) ReadHistory(
	identifier string,
	_ int,
) (session.History, error) {
	if _, okay := c.screens[identifier]; !okay {
		return session.History{}, fmt.Errorf(
			"session %s not found",
			identifier,
		)
	}

	return session.History{Identifier: identifier}, nil
}
