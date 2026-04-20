package tool

import (
	"context"
	"github.com/funtimecoding/go-library/pkg/generative/mark/response"
	"github.com/mark3labs/mcp-go/mcp"
)

type deleteProjectVariableArguments struct {
	Project string `json:"project"`
	Key     string `json:"key"`
}

func (t *Tool) DeleteProjectVariable(
	_ context.Context,
	_ mcp.CallToolRequest,
	a deleteProjectVariableArguments,
) (*mcp.CallToolResult, error) {
	if a.Project == "" {
		return response.Fail("project is required")
	}

	if a.Key == "" {
		return response.Fail("key is required")
	}

	_, e := t.client.ProjectVariables.RemoveVariable(
		a.Project,
		a.Key,
		nil,
	)

	if e != nil {
		return response.Fail("delete project variable: %v", e)
	}

	return mcp.NewToolResultText("deleted"), nil
}
