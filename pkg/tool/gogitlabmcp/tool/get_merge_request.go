package tool

import (
	"context"
	"github.com/funtimecoding/go-library/pkg/generative/mark/response"
	"github.com/mark3labs/mcp-go/mcp"
)

type getMergeRequestArguments struct {
	Project      string `json:"project"`
	MergeRequest int64  `json:"merge_request"`
}

func (t *Tool) GetMergeRequest(
	_ context.Context,
	_ mcp.CallToolRequest,
	a getMergeRequestArguments,
) (*mcp.CallToolResult, error) {
	if a.Project == "" {
		return response.Fail("project is required")
	}

	if a.MergeRequest == 0 {
		return response.Fail("merge_request is required")
	}

	v, _, e := t.client.MergeRequests.GetMergeRequest(
		a.Project,
		a.MergeRequest,
		nil,
	)

	if e != nil {
		return response.Fail("get merge request: %v", e)
	}

	return response.SuccessAny(v)
}
