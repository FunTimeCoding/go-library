package tool

import (
	"context"
	"github.com/funtimecoding/go-library/pkg/generative/mark/response"
	"github.com/mark3labs/mcp-go/mcp"
)

type openFileArguments struct {
	FilePath string `json:"file_path"`
}

func (t *Tool) OpenFile(
	_ context.Context,
	_ mcp.CallToolRequest,
	arguments openFileArguments,
) (*mcp.CallToolResult, error) {
	if arguments.FilePath == "" {
		return response.Fail("file_path is required")
	}

	v, e := t.client.OpenFile(arguments.FilePath)

	if e != nil {
		return response.Fail("open file: %v", e)
	}

	return response.SuccessAny(v)
}
