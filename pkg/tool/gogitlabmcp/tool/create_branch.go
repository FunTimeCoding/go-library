package tool

import (
	"context"
	"github.com/funtimecoding/go-library/pkg/generative/mark/response"
	"github.com/mark3labs/mcp-go/mcp"
	"gitlab.com/gitlab-org/api/client-go/v2"
)

type createBranchArguments struct {
	Project   string `json:"project"`
	Branch    string `json:"branch"`
	Reference string `json:"reference"`
}

func (t *Tool) CreateBranch(
	_ context.Context,
	_ mcp.CallToolRequest,
	a createBranchArguments,
) (*mcp.CallToolResult, error) {
	if a.Project == "" {
		return response.Fail("project is required")
	}

	if a.Branch == "" {
		return response.Fail("branch is required")
	}

	if a.Reference == "" {
		return response.Fail("reference is required")
	}

	v, _, e := t.client.Branches.CreateBranch(
		a.Project,
		&gitlab.CreateBranchOptions{
			Branch: &a.Branch,
			Ref:    &a.Reference,
		},
	)

	if e != nil {
		return response.Fail("create branch: %v", e)
	}

	return response.SuccessAny(v)
}
