package tool

import (
	"context"
	"github.com/funtimecoding/go-library/pkg/generative/mark/request"
	"github.com/funtimecoding/go-library/pkg/generative/mark/response"
	"github.com/mark3labs/mcp-go/mcp"
)

type editViewArguments struct {
	ViewIdentifier request.Integer `json:"view_id"`
	OldString      string          `json:"old_string"`
	NewString      string          `json:"new_string"`
	ReplaceAll     request.Boolean `json:"replace_all"`
}

func (t *Tool) EditView(
	_ context.Context,
	_ mcp.CallToolRequest,
	arguments editViewArguments,
) (*mcp.CallToolResult, error) {
	if arguments.OldString == "" {
		return response.Fail("old_string is required")
	}

	v, e := t.client.EditView(
		int(arguments.ViewIdentifier),
		arguments.OldString,
		arguments.NewString,
		bool(arguments.ReplaceAll),
	)

	if e != nil {
		return response.Fail("edit view: %v", e)
	}

	return response.SuccessAny(v)
}
