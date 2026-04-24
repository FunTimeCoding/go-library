package tool

import (
	"context"
	"github.com/funtimecoding/go-library/pkg/generative/mark/response"
	"github.com/funtimecoding/go-library/pkg/tool/gofirefoxmcp/argument"
	"github.com/mark3labs/mcp-go/mcp"
)

func (t *Tool) CloseTab(
	_ context.Context,
	_ mcp.CallToolRequest,
	a argument.CloseTab,
) (*mcp.CallToolResult, error) {
	e := t.client.CloseTab(int(a.TabIdentifier))

	if e != nil {
		return response.Fail("close tab: %v", e)
	}

	return response.Success("closed")
}
