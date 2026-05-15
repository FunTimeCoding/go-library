package model_context_client

import (
	"fmt"
	"github.com/modelcontextprotocol/go-sdk/mcp"
)

func (c *Client) CallToolError(
	name string,
	arguments map[string]any,
) (string, error) {
	result, e := c.CallToolRaw(name, arguments)

	if e != nil {
		return "", e
	}

	if !result.IsError {
		return "", fmt.Errorf("expected tool error, got success")
	}

	return result.Content[0].(*mcp.TextContent).Text, nil
}
