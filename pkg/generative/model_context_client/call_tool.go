package model_context_client

import (
	"fmt"
	"github.com/modelcontextprotocol/go-sdk/mcp"
)

func (c *Client) CallTool(
	name string,
	arguments map[string]any,
) (string, error) {
	result, e := c.CallToolRaw(name, arguments)

	if e != nil {
		return "", e
	}

	text := result.Content[0].(*mcp.TextContent).Text

	if result.IsError {
		return "", fmt.Errorf("tool error: %s", text)
	}

	return text, nil
}
