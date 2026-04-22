package tool

import (
	"context"
	"fmt"
	"github.com/funtimecoding/go-library/pkg/generative/mark/response"
	"github.com/mark3labs/mcp-go/mcp"
)

type getThreadRepliesArguments struct {
	PostID string `json:"post_id"`
}

func (t *Tool) GetThreadReplies(
	_ context.Context,
	_ mcp.CallToolRequest,
	a getThreadRepliesArguments,
) (result *mcp.CallToolResult, e error) {
	defer func() {
		if r := recover(); r != nil {
			result = mcp.NewToolResultError(fmt.Sprintf("%v", r))
			e = nil
		}
	}()

	if a.PostID == "" {
		return response.Fail("post_id is required")
	}

	parent := t.client.FindPost(a.PostID)
	replies := t.client.Thread(parent)
	t.client.Enrich(replies)
	type row struct {
		ID       string   `json:"id"`
		Username string   `json:"username"`
		Message  string   `json:"message"`
		CreateAt string   `json:"create_at"`
		FileIds  []string `json:"file_ids,omitempty"`
	}
	rows := make([]row, len(replies))

	for i, r := range replies {
		rows[i] = row{
			ID:       r.Raw.Id,
			Message:  r.Message,
			CreateAt: formatTime(r.Create),
		}

		if r.User != nil {
			rows[i].Username = r.User.Username
		}

		if len(r.Raw.FileIds) > 0 {
			rows[i].FileIds = r.Raw.FileIds
		}
	}

	return response.SuccessAny(
		map[string]any{
			"post_id": a.PostID,
			"replies": rows,
		},
	)
}
