package tool

import (
	"context"
	"github.com/funtimecoding/go-library/pkg/generative/mark/response"
	"github.com/funtimecoding/go-library/pkg/tool/gosublmcp/argument"
	"github.com/mark3labs/mcp-go/mcp"
)

func (t *Tool) CreateView(
	_ context.Context,
	_ mcp.CallToolRequest,
	a argument.CreateView,
) (*mcp.CallToolResult, error) {
	v, e := t.client.CreateView(a.Title, a.Content, a.Syntax)

	if e != nil {
		return response.Fail("create view: %v", e)
	}

	return response.SuccessAny(v)
}
