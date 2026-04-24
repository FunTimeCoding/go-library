package tool

import (
	"context"
	"github.com/funtimecoding/go-library/pkg/generative/mark/response"
	"github.com/funtimecoding/go-library/pkg/tool/gofirefoxmcp/argument"
	"github.com/mark3labs/mcp-go/mcp"
)

func (t *Tool) GroupTabs(
	_ context.Context,
	_ mcp.CallToolRequest,
	a argument.GroupTabs,
) (*mcp.CallToolResult, error) {
	if len(a.TabIdentifiers) == 0 {
		return response.Fail("tab_ids is required")
	}

	groupIdentifier, e := t.client.GroupTabs(
		a.TabIdentifiers,
		int(a.GroupIdentifier),
		a.Title,
		a.Color,
	)

	if e != nil {
		return response.Fail("group tabs: %v", e)
	}

	return response.Success("grouped into group %d", groupIdentifier)
}
