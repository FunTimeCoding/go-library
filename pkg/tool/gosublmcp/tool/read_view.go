package tool

import (
	"context"
	"github.com/funtimecoding/go-library/pkg/generative/mark/response"
	"github.com/funtimecoding/go-library/pkg/tool/gosublmcp/argument"
	"github.com/mark3labs/mcp-go/mcp"
)

func (t *Tool) ReadView(
	_ context.Context,
	_ mcp.CallToolRequest,
	a argument.ReadView,
) (*mcp.CallToolResult, error) {
	v, e := t.client.View(int(a.ViewIdentifier))

	if e != nil {
		return response.Fail("read view: %v", e)
	}

	return response.SuccessAny(v)
}
