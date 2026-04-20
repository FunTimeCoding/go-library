package tool

import (
	"context"
	"github.com/funtimecoding/go-library/pkg/generative/mark/response"
	"github.com/mark3labs/mcp-go/mcp"
	"gitlab.com/gitlab-org/api/client-go/v2"
)

type createMergeRequestNoteArguments struct {
	Project      string `json:"project"`
	MergeRequest int64  `json:"merge_request"`
	Body         string `json:"body"`
}

func (t *Tool) CreateMergeRequestNote(
	_ context.Context,
	_ mcp.CallToolRequest,
	a createMergeRequestNoteArguments,
) (*mcp.CallToolResult, error) {
	if a.Project == "" {
		return response.Fail("project is required")
	}

	if a.MergeRequest == 0 {
		return response.Fail("merge_request is required")
	}

	if a.Body == "" {
		return response.Fail("body is required")
	}

	v, _, e := t.client.Notes.CreateMergeRequestNote(
		a.Project,
		a.MergeRequest,
		&gitlab.CreateMergeRequestNoteOptions{Body: &a.Body},
	)

	if e != nil {
		return response.Fail("create merge request note: %v", e)
	}

	return response.SuccessAny(v)
}
