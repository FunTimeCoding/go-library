package mock_client

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/iterm/screen"
	"github.com/funtimecoding/go-library/pkg/iterm/session"
)

func (c *Client) CreateTab() (*session.Session, error) {
	id := fmt.Sprintf("session-%d", c.nextID)
	c.nextID++
	s := session.Stub()
	s.Identifier = id
	s.TabId = fmt.Sprintf("tab-%d", c.nextID-1)
	c.sessions = append(c.sessions, s)
	scr := screen.Stub()
	scr.Identifier = id
	c.screens[id] = scr

	return s, nil
}
