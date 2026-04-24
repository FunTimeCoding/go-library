package tool

import (
	"context"
	"github.com/funtimecoding/go-library/pkg/generative/mark/response"
	"github.com/funtimecoding/go-library/pkg/tool/gogitlabmcp/argument"
	"github.com/mark3labs/mcp-go/mcp"
)

func (t *Tool) GetCommit(
	_ context.Context,
	_ mcp.CallToolRequest,
	a argument.GetCommit,
) (*mcp.CallToolResult, error) {
	if a.Project == "" {
		return response.Fail("project is required")
	}

	if a.Sha == "" {
		return response.Fail("sha is required")
	}

	v, _, e := t.client.Commits.GetCommit(a.Project, a.Sha, nil)

	if e != nil {
		return response.Fail("get commit: %v", e)
	}

	return response.SuccessAny(v)
}
