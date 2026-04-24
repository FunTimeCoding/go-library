package tool

import (
	"context"
	"github.com/funtimecoding/go-library/pkg/generative/mark/response"
	"github.com/funtimecoding/go-library/pkg/tool/gogitlabmcp/argument"
	"github.com/mark3labs/mcp-go/mcp"
	"gitlab.com/gitlab-org/api/client-go/v2"
)

func (t *Tool) GetMergeRequestDiffs(
	_ context.Context,
	_ mcp.CallToolRequest,
	a argument.GetMergeRequestDiffs,
) (*mcp.CallToolResult, error) {
	if a.Project == "" {
		return response.Fail("project is required")
	}

	if a.MergeRequest == 0 {
		return response.Fail("merge_request is required")
	}

	v, _, e := t.client.MergeRequests.ListMergeRequestDiffs(
		a.Project,
		a.MergeRequest,
		&gitlab.ListMergeRequestDiffsOptions{
			ListOptions: gitlab.ListOptions{PerPage: 100},
		},
	)

	if e != nil {
		return response.Fail("get merge request diffs: %v", e)
	}

	return response.SuccessAny(v)
}
