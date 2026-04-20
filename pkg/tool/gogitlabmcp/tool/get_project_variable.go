package tool

import (
	"context"
	"github.com/funtimecoding/go-library/pkg/generative/mark/response"
	"github.com/mark3labs/mcp-go/mcp"
)

type getProjectVariableArguments struct {
	Project string `json:"project"`
	Key     string `json:"key"`
}

func (t *Tool) GetProjectVariable(
	_ context.Context,
	_ mcp.CallToolRequest,
	a getProjectVariableArguments,
) (*mcp.CallToolResult, error) {
	if a.Project == "" {
		return response.Fail("project is required")
	}

	if a.Key == "" {
		return response.Fail("key is required")
	}

	v, _, e := t.client.ProjectVariables.GetVariable(
		a.Project,
		a.Key,
		nil,
	)

	if e != nil {
		return response.Fail("get project variable: %v", e)
	}

	return response.SuccessAny(v)
}
