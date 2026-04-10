package tool

import (
	"context"
	"github.com/funtimecoding/go-library/pkg/generative/mark/response"
	"github.com/mark3labs/mcp-go/mcp"
)

type createTabArguments struct {
	Locator string `json:"url"`
}

func (t *Tool) CreateTab(
	_ context.Context,
	_ mcp.CallToolRequest,
	arguments createTabArguments,
) (*mcp.CallToolResult, error) {
	v, e := t.client.CreateTab(arguments.Locator)

	if e != nil {
		return response.Fail("create tab: %v", e)
	}

	return response.SuccessAny(v)
}
