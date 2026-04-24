package tool

import (
	"context"
	"github.com/funtimecoding/go-library/pkg/generative/mark/response"
	"github.com/funtimecoding/go-library/pkg/tool/gosublmcp/argument"
	"github.com/mark3labs/mcp-go/mcp"
)

func (t *Tool) OpenFile(
	_ context.Context,
	_ mcp.CallToolRequest,
	a argument.OpenFile,
) (*mcp.CallToolResult, error) {
	if a.FilePath == "" {
		return response.Fail("file_path is required")
	}

	v, e := t.client.OpenFile(a.FilePath)

	if e != nil {
		return response.Fail("open file: %v", e)
	}

	return response.SuccessAny(v)
}
