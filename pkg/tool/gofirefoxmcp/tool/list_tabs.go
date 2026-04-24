package tool

import (
	"context"
	"github.com/funtimecoding/go-library/pkg/generative/mark/response"
	"github.com/funtimecoding/go-library/pkg/tool/gofirefoxmcp/argument"
	"github.com/mark3labs/mcp-go/mcp"
)

func (t *Tool) ListTabs(
	_ context.Context,
	_ mcp.CallToolRequest,
	_ argument.ListTabs,
) (*mcp.CallToolResult, error) {
	v, e := t.client.Tabs()

	if e != nil {
		return response.Fail("list tabs: %v", e)
	}

	return response.SuccessAny(v)
}
