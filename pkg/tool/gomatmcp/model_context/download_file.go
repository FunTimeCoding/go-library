package model_context

import (
	"context"
	"github.com/funtimecoding/go-library/pkg/generative/mark/response"
	"github.com/funtimecoding/go-library/pkg/tool/gomatmcp/constant"
	"github.com/funtimecoding/go-library/pkg/tool/gomatmcp/model_context/argument"
	"github.com/mark3labs/mcp-go/mcp"
	"os"
	"path/filepath"
)

func (s *Server) DownloadFile(
	_ context.Context,
	_ mcp.CallToolRequest,
	a argument.DownloadFile,
) (*mcp.CallToolResult, error) {
	if a.FileIdentifier == "" {
		return response.Fail("file_id is required")
	}

	i, _, e := s.client.Nested().GetFileInfo(
		s.client.Context(),
		a.FileIdentifier,
	)

	if e != nil {
		return s.captureFail(e, constant.Unreachable)
	}

	b, _, e := s.client.Nested().GetFile(
		s.client.Context(),
		a.FileIdentifier,
	)

	if e != nil {
		return s.captureFail(e, constant.Unreachable)
	}

	path := a.Path

	if path == "" {
		path = filepath.Join(os.TempDir(), i.Name)
	}

	e = os.WriteFile(path, b, 0644)

	if e != nil {
		return s.captureFail(e, "file write failed")
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
