package tool

import (
	"context"
	"github.com/funtimecoding/go-library/pkg/generative/mark/response"
	"github.com/mark3labs/mcp-go/mcp"
	"gitlab.com/gitlab-org/api/client-go/v2"
)

type getRepositoryTreeArguments struct {
	Project   string `json:"project"`
	Path      string `json:"path"`
	Reference string `json:"reference"`
	Recursive bool   `json:"recursive"`
}

func (t *Tool) GetRepositoryTree(
	_ context.Context,
	_ mcp.CallToolRequest,
	a getRepositoryTreeArguments,
) (*mcp.CallToolResult, error) {
	if a.Project == "" {
		return response.Fail("project is required")
	}

	o := &gitlab.ListTreeOptions{
		ListOptions: gitlab.ListOptions{PerPage: 100},
	}

	if a.Path != "" {
		o.Path = &a.Path
	}

	if a.Reference != "" {
		o.Ref = &a.Reference
	}

	if a.Recursive {
		o.Recursive = &a.Recursive
	}

	v, _, e := t.client.Repositories.ListTree(a.Project, o)

	if e != nil {
		return response.Fail("get repository tree: %v", e)
	}

	return response.SuccessAny(v)
}
