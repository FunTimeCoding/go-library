package tool

import (
	"context"
	"github.com/funtimecoding/go-library/pkg/generative/mark/response"
	"github.com/mark3labs/mcp-go/mcp"
)

type createViewArguments struct {
	Title   string `json:"title"`
	Content string `json:"content"`
	Syntax  string `json:"syntax"`
}

func (t *Tool) CreateView(
	_ context.Context,
	_ mcp.CallToolRequest,
	arguments createViewArguments,
) (*mcp.CallToolResult, error) {
	v, e := t.client.CreateView(
		arguments.Title,
		arguments.Content,
		arguments.Syntax,
	)

	if e != nil {
		return response.Fail("create view: %v", e)
	}

	return response.SuccessAny(v)
}
