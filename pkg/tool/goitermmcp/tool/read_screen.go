package tool

import (
	"context"
	"github.com/funtimecoding/go-library/pkg/generative/mark/response"
	"github.com/mark3labs/mcp-go/mcp"
)

type readScreenArguments struct {
	SessionIdentifier string `json:"session_id"`
}

func (t *Tool) ReadScreen(
	_ context.Context,
	_ mcp.CallToolRequest,
	arguments readScreenArguments,
) (*mcp.CallToolResult, error) {
	v, e := t.client.ReadScreen(arguments.SessionIdentifier)

	if e != nil {
		return response.Fail("read screen: %v", e)
	}

	return response.SuccessAny(v)
}
