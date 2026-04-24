package tool

import (
	"context"
	"github.com/funtimecoding/go-library/pkg/generative/mark/response"
	"github.com/funtimecoding/go-library/pkg/tool/gogitlabmcp/argument"
	"github.com/mark3labs/mcp-go/mcp"
)

func (t *Tool) GetProjectVariable(
	_ context.Context,
	_ mcp.CallToolRequest,
	a argument.GetProjectVariable,
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
