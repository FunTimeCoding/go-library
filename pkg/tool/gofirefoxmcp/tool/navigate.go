package tool

import (
	"context"
	"github.com/funtimecoding/go-library/pkg/generative/mark/request"
	"github.com/funtimecoding/go-library/pkg/generative/mark/response"
	"github.com/mark3labs/mcp-go/mcp"
)

type navigateArguments struct {
	TabIdentifier request.Integer `json:"tab_id"`
	Locator       string          `json:"url"`
}

func (t *Tool) Navigate(
	_ context.Context,
	_ mcp.CallToolRequest,
	arguments navigateArguments,
) (*mcp.CallToolResult, error) {
	if arguments.Locator == "" {
		return response.Fail("url is required")
	}

	e := t.client.Navigate(int(arguments.TabIdentifier), arguments.Locator)

	if e != nil {
		return response.Fail("navigate: %v", e)
	}

	return response.Success("navigated to %s", arguments.Locator)
}
