package tool

import (
	"context"
	"github.com/funtimecoding/go-library/pkg/generative/mark/response"
	"github.com/funtimecoding/go-library/pkg/tool/gofirefoxmcp/argument"
	"github.com/mark3labs/mcp-go/mcp"
)

func (t *Tool) ReadTab(
	_ context.Context,
	_ mcp.CallToolRequest,
	a argument.ReadTab,
) (*mcp.CallToolResult, error) {
	v, e := t.client.ReadTab(
		int(a.TabIdentifier),
		bool(a.Raw),
	)

	if e != nil {
		return response.Fail("read tab: %v", e)
	}

	return response.SuccessAny(v)
}
