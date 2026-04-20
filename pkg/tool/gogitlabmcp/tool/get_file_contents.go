package tool

import (
	"context"
	"github.com/funtimecoding/go-library/pkg/generative/mark/response"
	"github.com/mark3labs/mcp-go/mcp"
	"gitlab.com/gitlab-org/api/client-go/v2"
)

type getFileContentsArguments struct {
	Project   string `json:"project"`
	Path      string `json:"path"`
	Reference string `json:"reference"`
}

func (t *Tool) GetFileContents(
	_ context.Context,
	_ mcp.CallToolRequest,
	a getFileContentsArguments,
) (*mcp.CallToolResult, error) {
	if a.Project == "" {
		return response.Fail("project is required")
	}

	if a.Path == "" {
		return response.Fail("path is required")
	}

	o := &gitlab.GetFileOptions{}

	if a.Reference != "" {
		o.Ref = &a.Reference
	}

	v, _, e := t.client.RepositoryFiles.GetFile(a.Project, a.Path, o)

	if e != nil {
		return response.Fail("get file contents: %v", e)
	}

	return response.SuccessAny(v)
}
