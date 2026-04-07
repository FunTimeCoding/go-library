package tool

import (
	"context"
	"fmt"
	"github.com/funtimecoding/go-library/pkg/generative/mark/response"
	"github.com/mark3labs/mcp-go/mcp"
	"time"
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
	type row struct {
		ID       string `json:"id"`
		UserID   string `json:"user_id"`
		Message  string `json:"message"`
		CreateAt string `json:"create_at"`
	}
	rows := make([]row, len(replies))

	for i, r := range replies {
		rows[i] = row{
			ID:       r.Raw.Id,
			UserID:   r.Raw.UserId,
			Message:  r.Message,
			CreateAt: r.Create.Format(time.RFC3339),
		}
	}

	return response.SuccessAny(
		map[string]any{
			"post_id": a.PostID,
			"replies": rows,
		},
	)
}
