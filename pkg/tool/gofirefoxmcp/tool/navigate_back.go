package tool

import (
	"context"
	"github.com/funtimecoding/go-library/pkg/generative/mark/response"
	"github.com/funtimecoding/go-library/pkg/tool/gofirefoxmcp/argument"
	"github.com/mark3labs/mcp-go/mcp"
)

func (t *Tool) NavigateBack(
	_ context.Context,
	_ mcp.CallToolRequest,
	a argument.NavigateBack,
) (*mcp.CallToolResult, error) {
	e := t.client.NavigateBack(int(a.TabIdentifier))

	if e != nil {
		return response.Fail("navigate back: %v", e)
	}

	return response.Success("navigated back")
}
