package tool

import (
	"context"
	"github.com/funtimecoding/go-library/pkg/generative/mark/response"
	"github.com/funtimecoding/go-library/pkg/tool/gofirefoxmcp/argument"
	"github.com/mark3labs/mcp-go/mcp"
)

func (t *Tool) UpdateGroup(
	_ context.Context,
	_ mcp.CallToolRequest,
	a argument.UpdateGroup,
) (*mcp.CallToolResult, error) {
	e := t.client.UpdateGroup(
		int(a.GroupIdentifier),
		a.Title,
		a.Color,
		nil,
	)

	if e != nil {
		return response.Fail("update group: %v", e)
	}

	return response.Success("group updated")
}
