package tool

import (
	"context"
	"github.com/funtimecoding/go-library/pkg/generative/mark/response"
	"github.com/funtimecoding/go-library/pkg/tool/gogitlabmcp/argument"
	"github.com/mark3labs/mcp-go/mcp"
)

func (t *Tool) GetProject(
	_ context.Context,
	_ mcp.CallToolRequest,
	a argument.GetProject,
) (*mcp.CallToolResult, error) {
	if a.Project == "" {
		return response.Fail("project is required")
	}

	v, _, e := t.client.Projects.GetProject(a.Project, nil)

	if e != nil {
		return response.Fail("get project: %v", e)
	}

	return response.SuccessAny(v)
}
