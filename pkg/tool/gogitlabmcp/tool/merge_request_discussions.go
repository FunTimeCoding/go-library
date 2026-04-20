package tool

import (
	"context"
	"github.com/funtimecoding/go-library/pkg/generative/mark/response"
	"github.com/mark3labs/mcp-go/mcp"
	"gitlab.com/gitlab-org/api/client-go/v2"
)

type mergeRequestDiscussionsArguments struct {
	Project      string `json:"project"`
	MergeRequest int64  `json:"merge_request"`
}

func (t *Tool) MergeRequestDiscussions(
	_ context.Context,
	_ mcp.CallToolRequest,
	a mergeRequestDiscussionsArguments,
) (*mcp.CallToolResult, error) {
	if a.Project == "" {
		return response.Fail("project is required")
	}

	if a.MergeRequest == 0 {
		return response.Fail("merge_request is required")
	}

	v, _, e := t.client.Discussions.ListMergeRequestDiscussions(
		a.Project,
		a.MergeRequest,
		&gitlab.ListMergeRequestDiscussionsOptions{
			ListOptions: gitlab.ListOptions{PerPage: 100},
		},
	)

	if e != nil {
		return response.Fail("list merge request discussions: %v", e)
	}

	return response.SuccessAny(v)
}
