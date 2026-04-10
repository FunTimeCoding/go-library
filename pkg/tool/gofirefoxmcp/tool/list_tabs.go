package tool

import (
	"context"
	"github.com/funtimecoding/go-library/pkg/generative/mark/response"
	"github.com/mark3labs/mcp-go/mcp"
)

type listTabsArguments struct{}

func (t *Tool) ListTabs(
	_ context.Context,
	_ mcp.CallToolRequest,
	_ listTabsArguments,
) (*mcp.CallToolResult, error) {
	v, e := t.client.Tabs()

	if e != nil {
		return response.Fail("list tabs: %v", e)
	}

	return response.SuccessAny(v)
}
