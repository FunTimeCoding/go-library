package tool

import (
	"context"
	"github.com/funtimecoding/go-library/pkg/generative/mark/response"
	"github.com/mark3labs/mcp-go/mcp"
	"gitlab.com/gitlab-org/api/client-go/v2"
)

type updateProjectVariableArguments struct {
	Project   string `json:"project"`
	Key       string `json:"key"`
	Value     string `json:"value"`
	Protected bool   `json:"protected"`
	Masked    bool   `json:"masked"`
}

func (t *Tool) UpdateProjectVariable(
	_ context.Context,
	_ mcp.CallToolRequest,
	a updateProjectVariableArguments,
) (*mcp.CallToolResult, error) {
	if a.Project == "" {
		return response.Fail("project is required")
	}

	if a.Key == "" {
		return response.Fail("key is required")
	}

	if a.Value == "" {
		return response.Fail("value is required")
	}

	v, _, e := t.client.ProjectVariables.UpdateVariable(
		a.Project,
		a.Key,
		&gitlab.UpdateProjectVariableOptions{
			Value:     &a.Value,
			Protected: &a.Protected,
			Masked:    &a.Masked,
		},
	)

	if e != nil {
		return response.Fail("update project variable: %v", e)
	}

	return response.SuccessAny(v)
}
