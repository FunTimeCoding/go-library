package mock_client

import "github.com/funtimecoding/go-library/pkg/generative/anthropic/claude/message"

func (c *Client) Messages(sessionIdentifier string) []message.Message {
	return nil
}
