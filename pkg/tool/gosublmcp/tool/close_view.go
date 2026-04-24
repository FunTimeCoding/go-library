package tool

import (
	"context"
	"github.com/funtimecoding/go-library/pkg/generative/mark/response"
	"github.com/funtimecoding/go-library/pkg/tool/gosublmcp/argument"
	"github.com/mark3labs/mcp-go/mcp"
)

func (t *Tool) CloseView(
	_ context.Context,
	_ mcp.CallToolRequest,
	a argument.CloseView,
) (*mcp.CallToolResult, error) {
	e := t.client.CloseView(int(a.ViewIdentifier))

	if e != nil {
		return response.Fail("close view: %v", e)
	}

	return response.Success("closed")
}
