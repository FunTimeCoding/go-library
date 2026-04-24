package tool

import (
	"context"
	"github.com/funtimecoding/go-library/pkg/generative/mark/response"
	"github.com/funtimecoding/go-library/pkg/tool/goitermmcp/argument"
	"github.com/mark3labs/mcp-go/mcp"
)

func (t *Tool) SetTabColor(
	_ context.Context,
	_ mcp.CallToolRequest,
	a argument.SetTabColor,
) (*mcp.CallToolResult, error) {
	e := t.client.SetTabColor(
		a.SessionIdentifier,
		int(a.Red),
		int(a.Green),
		int(a.Blue),
	)

	if e != nil {
		return response.Fail("set tab color: %v", e)
	}

	return response.Success("color set")
}
