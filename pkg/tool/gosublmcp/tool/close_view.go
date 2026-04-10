package tool

import (
	"context"
	"github.com/funtimecoding/go-library/pkg/generative/mark/request"
	"github.com/funtimecoding/go-library/pkg/generative/mark/response"
	"github.com/mark3labs/mcp-go/mcp"
)

type closeViewArguments struct {
	ViewIdentifier request.Integer `json:"view_id"`
}

func (t *Tool) CloseView(
	_ context.Context,
	_ mcp.CallToolRequest,
	arguments closeViewArguments,
) (*mcp.CallToolResult, error) {
	e := t.client.CloseView(int(arguments.ViewIdentifier))

	if e != nil {
		return response.Fail("close view: %v", e)
	}

	return response.Success("closed")
}
