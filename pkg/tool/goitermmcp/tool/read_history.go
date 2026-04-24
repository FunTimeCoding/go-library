package tool

import (
	"context"
	"github.com/funtimecoding/go-library/pkg/generative/mark/response"
	"github.com/funtimecoding/go-library/pkg/tool/goitermmcp/argument"
	"github.com/mark3labs/mcp-go/mcp"
)

func (t *Tool) ReadHistory(
	_ context.Context,
	_ mcp.CallToolRequest,
	a argument.ReadHistory,
) (*mcp.CallToolResult, error) {
	count := int(a.Count)

	if count <= 0 {
		count = 50
	}

	v, e := t.client.ReadHistory(a.SessionIdentifier, count)

	if e != nil {
		return response.Fail("read history: %v", e)
	}

	return response.SuccessAny(v)
}
