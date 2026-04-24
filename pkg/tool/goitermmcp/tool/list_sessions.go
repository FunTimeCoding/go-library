package tool

import (
	"context"
	"github.com/funtimecoding/go-library/pkg/generative/mark/response"
	"github.com/funtimecoding/go-library/pkg/tool/goitermmcp/argument"
	"github.com/mark3labs/mcp-go/mcp"
)

func (t *Tool) ListSessions(
	_ context.Context,
	_ mcp.CallToolRequest,
	_ argument.ListSessions,
) (*mcp.CallToolResult, error) {
	v, e := t.client.Sessions()

	if e != nil {
		return response.Fail("list sessions: %v", e)
	}

	return response.SuccessAny(v)
}
