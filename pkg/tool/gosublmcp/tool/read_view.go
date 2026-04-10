package tool

import (
	"context"
	"github.com/funtimecoding/go-library/pkg/generative/mark/request"
	"github.com/funtimecoding/go-library/pkg/generative/mark/response"
	"github.com/mark3labs/mcp-go/mcp"
)

type readViewArguments struct {
	ViewIdentifier request.Integer `json:"view_id"`
}

func (t *Tool) ReadView(
	_ context.Context,
	_ mcp.CallToolRequest,
	arguments readViewArguments,
) (*mcp.CallToolResult, error) {
	v, e := t.client.View(int(arguments.ViewIdentifier))

	if e != nil {
		return response.Fail("read view: %v", e)
	}

	return response.SuccessAny(v)
}
