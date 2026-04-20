package tool

import (
	"context"
	"github.com/funtimecoding/go-library/pkg/generative/mark/response"
	"github.com/mark3labs/mcp-go/mcp"
	"gitlab.com/gitlab-org/api/client-go/v2"
)

type listProjectVariablesArguments struct {
	Project string `json:"project"`
}

func (t *Tool) ListProjectVariables(
	_ context.Context,
	_ mcp.CallToolRequest,
	a listProjectVariablesArguments,
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
