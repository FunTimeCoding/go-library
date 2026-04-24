package tool

import (
	"context"
	"github.com/funtimecoding/go-library/pkg/generative/mark/response"
	"github.com/funtimecoding/go-library/pkg/tool/gogitlabmcp/argument"
	"github.com/mark3labs/mcp-go/mcp"
	"gitlab.com/gitlab-org/api/client-go/v2"
)

func (t *Tool) ListProjectVariables(
	_ context.Context,
	_ mcp.CallToolRequest,
	a argument.ListProjectVariables,
) (*mcp.CallToolResult, error) {
	if a.Project == "" {
		return response.Fail("project is required")
	}

	v, _, e := t.client.ProjectVariables.ListVariables(
		a.Project,
		&gitlab.ListProjectVariablesOptions{},
	)

	if e != nil {
		return response.Fail("list project variables: %v", e)
	}

	return response.SuccessAny(v)
}
