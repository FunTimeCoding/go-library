package tool

import (
	"context"
	"github.com/funtimecoding/go-library/pkg/generative/mark/request"
	"github.com/funtimecoding/go-library/pkg/generative/mark/response"
	"github.com/mark3labs/mcp-go/mcp"
)

type readTabArguments struct {
	TabIdentifier request.Integer `json:"tab_id"`
	Raw           request.Boolean `json:"raw"`
}

func (t *Tool) ReadTab(
	_ context.Context,
	_ mcp.CallToolRequest,
	arguments readTabArguments,
) (*mcp.CallToolResult, error) {
	v, e := t.client.ReadTab(int(arguments.TabIdentifier), bool(arguments.Raw))

	if e != nil {
		return response.Fail("read tab: %v", e)
	}

	return response.SuccessAny(v)
}
