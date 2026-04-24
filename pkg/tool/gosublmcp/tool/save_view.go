package tool

import (
	"context"
	"github.com/funtimecoding/go-library/pkg/generative/mark/response"
	"github.com/funtimecoding/go-library/pkg/tool/gosublmcp/argument"
	"github.com/mark3labs/mcp-go/mcp"
)

func (t *Tool) SaveView(
	_ context.Context,
	_ mcp.CallToolRequest,
	a argument.SaveView,
) (*mcp.CallToolResult, error) {
	e := t.client.SaveView(int(a.ViewIdentifier), a.FilePath)

	if e != nil {
		return response.Fail("save view: %v", e)
	}

	return response.Success("saved")
}
