package tool

import (
	"context"
	"github.com/funtimecoding/go-library/pkg/generative/mark/request"
	"github.com/funtimecoding/go-library/pkg/generative/mark/response"
	"github.com/mark3labs/mcp-go/mcp"
)

type groupTabsArguments struct {
	TabIdentifiers  []int  `json:"tab_ids"`
	GroupIdentifier request.Integer `json:"group_id"`
	Title           string `json:"title"`
	Color           string `json:"color"`
}

func (t *Tool) GroupTabs(
	_ context.Context,
	_ mcp.CallToolRequest,
	arguments groupTabsArguments,
) (*mcp.CallToolResult, error) {
	if len(arguments.TabIdentifiers) == 0 {
		return response.Fail("tab_ids is required")
	}

	groupIdentifier, e := t.client.GroupTabs(
		arguments.TabIdentifiers,
		int(arguments.GroupIdentifier),
		arguments.Title,
		arguments.Color,
	)

	if e != nil {
		return response.Fail("group tabs: %v", e)
	}

	return response.Success("grouped into group %d", groupIdentifier)
}
