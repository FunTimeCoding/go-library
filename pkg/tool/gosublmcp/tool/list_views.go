package tool

import (
	"context"
	"github.com/funtimecoding/go-library/pkg/generative/mark/response"
	"github.com/mark3labs/mcp-go/mcp"
)

type listViewsArguments struct{}

func (t *Tool) ListViews(
	_ context.Context,
	_ mcp.CallToolRequest,
	_ listViewsArguments,
) (*mcp.CallToolResult, error) {
	v, e := t.client.Views()

	if e != nil {
		return response.Fail("list views: %v", e)
	}

	return response.SuccessAny(v)
}
