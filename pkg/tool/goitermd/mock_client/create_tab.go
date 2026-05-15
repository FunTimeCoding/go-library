package mock_client

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/iterm/session"
)

func (c *Client) CreateTab() (session.Session, error) {
	id := fmt.Sprintf("session-%d", c.nextID)
	c.nextID++
	s := session.Session{
		Identifier: id,
		TabId:      fmt.Sprintf("tab-%d", c.nextID-1),
	}
	c.sessions = append(c.sessions, s)
	c.screens[id] = session.Screen{Identifier: id}

	return s, nil
}
