package mock_client

import "github.com/funtimecoding/go-library/pkg/generative/anthropic/claude"

func (c *Client) ToolContext(
	sessionIdentifier string,
	toolFilter string,
	surroundCount int,
) []claude.ToolContextResult {
	return nil
}
