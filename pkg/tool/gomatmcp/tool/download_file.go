package tool

import (
	"context"
	"fmt"
	"github.com/funtimecoding/go-library/pkg/generative/mark/response"
	"github.com/funtimecoding/go-library/pkg/tool/gomatmcp/argument"
	"github.com/mark3labs/mcp-go/mcp"
	"os"
	"path/filepath"
)

func (t *Tool) DownloadFile(
	_ context.Context,
	_ mcp.CallToolRequest,
	a argument.DownloadFile,
) (result *mcp.CallToolResult, e error) {
	defer func() {
		if r := recover(); r != nil {
			result = mcp.NewToolResultError(fmt.Sprintf("%v", r))
			e = nil
		}
	}()

	if a.FileIdentifier == "" {
		return response.Fail("file_id is required")
	}

	info, _, f := t.client.Nested().GetFileInfo(
		t.client.Context(),
		a.FileIdentifier,
	)

	if f != nil {
		return response.Fail("get file info failed: %v", f)
	}

	data, _, g := t.client.Nested().GetFile(
		t.client.Context(),
		a.FileIdentifier,
	)

	if g != nil {
		return response.Fail("download failed: %v", g)
	}

	path := a.Path

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
