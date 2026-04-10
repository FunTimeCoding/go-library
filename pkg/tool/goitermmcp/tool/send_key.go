package tool

import (
	"context"
	"github.com/funtimecoding/go-library/pkg/generative/mark/response"
	"github.com/mark3labs/mcp-go/mcp"
)

type sendKeyArguments struct {
	SessionIdentifier string `json:"session_id"`
	Key               string `json:"key"`
}

func (t *Tool) SendKey(
	_ context.Context,
	_ mcp.CallToolRequest,
	arguments sendKeyArguments,
) (*mcp.CallToolResult, error) {
	if arguments.Key == "" {
		return response.Fail("key is required")
	}

	e := t.client.SendKey(arguments.SessionIdentifier, arguments.Key)

	if e != nil {
		return response.Fail("send key: %v", e)
	}

	return response.Success("sent key: %s", arguments.Key)
}
