package tool

import (
	"context"
	"github.com/funtimecoding/go-library/pkg/generative/mark/request"
	"github.com/funtimecoding/go-library/pkg/generative/mark/response"
	"github.com/mark3labs/mcp-go/mcp"
)

type ungroupTabArguments struct {
	TabIdentifier request.Integer `json:"tab_id"`
}

func (t *Tool) UngroupTab(
	_ context.Context,
	_ mcp.CallToolRequest,
	arguments ungroupTabArguments,
) (*mcp.CallToolResult, error) {
	e := t.client.UngroupTab(int(arguments.TabIdentifier))

	if e != nil {
		return response.Fail("ungroup tab: %v", e)
	}

	return response.Success("ungrouped")
}
