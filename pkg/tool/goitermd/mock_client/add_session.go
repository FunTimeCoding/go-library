package mock_client

import "github.com/funtimecoding/go-library/pkg/iterm/session"

func (c *Client) AddSession(s session.Session) {
	c.sessions = append(c.sessions, s)
	c.screens[s.Identifier] = session.Screen{
		Identifier: s.Identifier,
	}
}
