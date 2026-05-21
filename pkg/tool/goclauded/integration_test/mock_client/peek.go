package mock_client

import "github.com/funtimecoding/go-library/pkg/generative/anthropic/claude/peek"

func (c *Client) Peek(sessionIdentifier string) *peek.Peek {
	if messages, okay := c.UserMessages[sessionIdentifier]; okay {
		return peek.New(messages)
	}

	return peek.Stub()
}
