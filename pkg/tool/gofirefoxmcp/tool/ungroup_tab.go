package tool

import (
	"context"
	"github.com/funtimecoding/go-library/pkg/generative/mark/response"
	"github.com/funtimecoding/go-library/pkg/tool/gofirefoxmcp/argument"
	"github.com/mark3labs/mcp-go/mcp"
)

func (t *Tool) UngroupTab(
	_ context.Context,
	_ mcp.CallToolRequest,
	a argument.UngroupTab,
) (*mcp.CallToolResult, error) {
	e := t.client.UngroupTab(int(a.TabIdentifier))

	if e != nil {
		return response.Fail("ungroup tab: %v", e)
	}

	return response.Success("ungrouped")
}
