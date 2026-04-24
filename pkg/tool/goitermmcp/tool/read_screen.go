package tool

import (
	"context"
	"github.com/funtimecoding/go-library/pkg/generative/mark/response"
	"github.com/funtimecoding/go-library/pkg/tool/goitermmcp/argument"
	"github.com/mark3labs/mcp-go/mcp"
)

func (t *Tool) ReadScreen(
	_ context.Context,
	_ mcp.CallToolRequest,
	a argument.ReadScreen,
) (*mcp.CallToolResult, error) {
	v, e := t.client.ReadScreen(a.SessionIdentifier)

	if e != nil {
		return response.Fail("read screen: %v", e)
	}

	return response.SuccessAny(v)
}
