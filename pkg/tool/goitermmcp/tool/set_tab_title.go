package tool

import (
	"context"
	"github.com/funtimecoding/go-library/pkg/generative/mark/response"
	"github.com/mark3labs/mcp-go/mcp"
)

type setTabTitleArguments struct {
	TabIdentifier string `json:"tab_id"`
	Title         string `json:"title"`
}

func (t *Tool) SetTabTitle(
	_ context.Context,
	_ mcp.CallToolRequest,
	arguments setTabTitleArguments,
) (*mcp.CallToolResult, error) {
	if arguments.Title == "" {
		return response.Fail("title is required")
	}

	e := t.client.SetTabTitle(arguments.TabIdentifier, arguments.Title)

	if e != nil {
		return response.Fail("set tab title: %v", e)
	}

	return response.Success("set title: %s", arguments.Title)
}
