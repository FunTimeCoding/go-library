package tool

import (
	"context"
	"github.com/funtimecoding/go-library/pkg/generative/mark/request"
	"github.com/funtimecoding/go-library/pkg/generative/mark/response"
	"github.com/mark3labs/mcp-go/mcp"
)

type closeTabArguments struct {
	TabIdentifier request.Integer `json:"tab_id"`
}

func (t *Tool) CloseTab(
	_ context.Context,
	_ mcp.CallToolRequest,
	arguments closeTabArguments,
) (*mcp.CallToolResult, error) {
	e := t.client.CloseTab(int(arguments.TabIdentifier))

	if e != nil {
		return response.Fail("close tab: %v", e)
	}

	return response.Success("closed")
}
