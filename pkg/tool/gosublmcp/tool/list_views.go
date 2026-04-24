package tool

import (
	"context"
	"github.com/funtimecoding/go-library/pkg/generative/mark/response"
	"github.com/funtimecoding/go-library/pkg/tool/gosublmcp/argument"
	"github.com/mark3labs/mcp-go/mcp"
)

func (t *Tool) ListViews(
	_ context.Context,
	_ mcp.CallToolRequest,
	_ argument.ListViews,
) (*mcp.CallToolResult, error) {
	v, e := t.client.Views()

	if e != nil {
		return response.Fail("list views: %v", e)
	}

	return response.SuccessAny(v)
}
