package tool

import (
	"context"
	"github.com/funtimecoding/go-library/pkg/generative/mark/response"
	"github.com/funtimecoding/go-library/pkg/tool/goitermmcp/argument"
	"github.com/mark3labs/mcp-go/mcp"
)

func (t *Tool) CreateTab(
	_ context.Context,
	_ mcp.CallToolRequest,
	_ argument.CreateTab,
) (*mcp.CallToolResult, error) {
	v, e := t.client.CreateTab()

	if e != nil {
		return response.Fail("create tab: %v", e)
	}

	return response.SuccessAny(v)
}
