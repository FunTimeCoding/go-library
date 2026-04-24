package tool

import (
	"context"
	"github.com/funtimecoding/go-library/pkg/generative/mark/response"
	"github.com/funtimecoding/go-library/pkg/tool/gosublmcp/argument"
	"github.com/mark3labs/mcp-go/mcp"
)

func (t *Tool) EditView(
	_ context.Context,
	_ mcp.CallToolRequest,
	a argument.EditView,
) (*mcp.CallToolResult, error) {
	if a.OldString == "" {
		return response.Fail("old_string is required")
	}

	v, e := t.client.EditView(
		int(a.ViewIdentifier),
		a.OldString,
		a.NewString,
		bool(a.ReplaceAll),
	)

	if e != nil {
		return response.Fail("edit view: %v", e)
	}

	return response.SuccessAny(v)
}
