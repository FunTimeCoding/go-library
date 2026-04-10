package tool

import (
	"context"
	"github.com/funtimecoding/go-library/pkg/generative/mark/request"
	"github.com/funtimecoding/go-library/pkg/generative/mark/response"
	"github.com/mark3labs/mcp-go/mcp"
)

type navigateBackArguments struct {
	TabIdentifier request.Integer `json:"tab_id"`
}

func (t *Tool) NavigateBack(
	_ context.Context,
	_ mcp.CallToolRequest,
	arguments navigateBackArguments,
) (*mcp.CallToolResult, error) {
	e := t.client.NavigateBack(int(arguments.TabIdentifier))

	if e != nil {
		return response.Fail("navigate back: %v", e)
	}

	return response.Success("navigated back")
}
