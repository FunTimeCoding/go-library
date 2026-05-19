package claude

import (
	"github.com/funtimecoding/go-library/pkg/generative/anthropic/claude/session"
	"strings"
)

func (c *Client) Resolve(query string) *session.Session {
	for _, s := range c.Sessions() {
		if s.Identifier == query {
			return s
		}

		if s.Slug == query {
			return s
		}

		if strings.HasPrefix(s.Identifier, query) {
			return s
		}
	}

	return session.Stub()
}
