package tool

import (
	"context"
	"github.com/funtimecoding/go-library/pkg/generative/mark/response"
	"github.com/funtimecoding/go-library/pkg/tool/gofirefoxmcp/argument"
	"github.com/mark3labs/mcp-go/mcp"
)

func (t *Tool) Navigate(
	_ context.Context,
	_ mcp.CallToolRequest,
	a argument.Navigate,
) (*mcp.CallToolResult, error) {
	if a.Locator == "" {
		return response.Fail("url is required")
	}

	e := t.client.Navigate(int(a.TabIdentifier), a.Locator)

	if e != nil {
		return response.Fail("navigate: %v", e)
	}

	return response.Success("navigated to %s", a.Locator)
}
