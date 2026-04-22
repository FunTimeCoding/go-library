package tool

import (
	"context"
	"fmt"
	"github.com/funtimecoding/go-library/pkg/generative/mark/response"
	"github.com/mark3labs/mcp-go/mcp"
	"os"
	"path/filepath"
)

type downloadFileArguments struct {
	FileIdentifier string `json:"file_id"`
	Path           string `json:"path"`
}

func (t *Tool) DownloadFile(
	_ context.Context,
	_ mcp.CallToolRequest,
	arguments downloadFileArguments,
) (result *mcp.CallToolResult, e error) {
	defer func() {
		if r := recover(); r != nil {
			result = mcp.NewToolResultError(fmt.Sprintf("%v", r))
			e = nil
		}
	}()

	if arguments.FileIdentifier == "" {
		return response.Fail("file_id is required")
	}

	info, _, f := t.client.Nested().GetFileInfo(
		t.client.Context(),
		arguments.FileIdentifier,
	)

	if f != nil {
		return response.Fail("get file info failed: %v", f)
	}

	data, _, g := t.client.Nested().GetFile(
		t.client.Context(),
		arguments.FileIdentifier,
	)

	if g != nil {
		return response.Fail("download failed: %v", g)
	}

	path := arguments.Path

	if path == "" {
		path = filepath.Join(os.TempDir(), info.Name)
	}

	h := os.WriteFile(path, data, 0644)

	if h != nil {
		return response.Fail("write failed: %v", h)
	}

	return response.SuccessAny(
		map[string]any{
			"file_id":   info.Id,
			"name":      info.Name,
			"mime_type": info.MimeType,
			"size":      info.Size,
			"path":      path,
		},
	)
}
