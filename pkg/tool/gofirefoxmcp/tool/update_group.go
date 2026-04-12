package tool

import (
	"context"
	"github.com/funtimecoding/go-library/pkg/generative/mark/request"
	"github.com/funtimecoding/go-library/pkg/generative/mark/response"
	"github.com/mark3labs/mcp-go/mcp"
)

type updateGroupArguments struct {
	GroupIdentifier request.Integer `json:"group_id"`
	Title           string          `json:"title"`
	Color           string          `json:"color"`
}

func (t *Tool) UpdateGroup(
	_ context.Context,
	_ mcp.CallToolRequest,
	arguments updateGroupArguments,
) (*mcp.CallToolResult, error) {
	e := t.client.UpdateGroup(
		int(arguments.GroupIdentifier),
		arguments.Title,
		arguments.Color,
		nil,
	)

	if e != nil {
		return response.Fail("update group: %v", e)
	}

	return response.Success("group updated")
}
