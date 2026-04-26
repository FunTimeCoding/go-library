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

	i, _, f := t.client.Nested().GetFileInfo(
		t.client.Context(),
		a.FileIdentifier,
	)

	if f != nil {
		return response.Fail("get file info failed: %v", f)
	}

	b, _, g := t.client.Nested().GetFile(
		t.client.Context(),
		a.FileIdentifier,
	)

	if g != nil {
		return response.Fail("download failed: %v", g)
	}

	path := a.Path

	if path == "" {
		path = filepath.Join(os.TempDir(), i.Name)
	}

	h := os.WriteFile(path, b, 0644)

	if h != nil {
		return response.Fail("write failed: %v", h)
	}

	return response.SuccessAny(
		map[string]any{
			"file_id":   i.Id,
			"name":      i.Name,
			"mime_type": i.MimeType,
			"size":      i.Size,
			"path":      path,
		},
	)
}
