package mock_client

import "github.com/funtimecoding/go-library/pkg/generative/anthropic/claude/tool_call"

func (c *Client) ToolCalls(sessionIdentifier string) []tool_call.Call {
	return nil
}
