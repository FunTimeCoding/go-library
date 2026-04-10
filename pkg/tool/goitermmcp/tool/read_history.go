package tool

import (
	"context"
	"github.com/funtimecoding/go-library/pkg/generative/mark/request"
	"github.com/funtimecoding/go-library/pkg/generative/mark/response"
	"github.com/mark3labs/mcp-go/mcp"
)

type readHistoryArguments struct {
	SessionIdentifier string `json:"session_id"`
	Count             request.Integer `json:"count"`
}

func (t *Tool) ReadHistory(
	_ context.Context,
	_ mcp.CallToolRequest,
	arguments readHistoryArguments,
) (*mcp.CallToolResult, error) {
	count := int(arguments.Count)

	if count <= 0 {
		count = 50
	}

	v, e := t.client.ReadHistory(arguments.SessionIdentifier, count)

	if e != nil {
		return response.Fail("read history: %v", e)
	}

	return response.SuccessAny(v)
}
