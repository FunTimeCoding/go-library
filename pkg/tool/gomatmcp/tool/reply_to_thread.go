package tool

import (
	"context"
	"fmt"
	"github.com/funtimecoding/go-library/pkg/generative/mark/response"
	"github.com/mark3labs/mcp-go/mcp"
)

type replyToThreadArguments struct {
	ChannelID string `json:"channel_id"`
	PostID    string `json:"post_id"`
	Message   string `json:"message"`
}

func (t *Tool) ReplyToThread(
	_ context.Context,
	_ mcp.CallToolRequest,
	a replyToThreadArguments,
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

	if a.Message == "" {
		return response.Fail("message is required")
	}

	p := t.client.Reply(
		t.client.Channel(a.ChannelID),
		t.client.FindPost(a.PostID),
		a.Message,
	)

	return response.SuccessAny(
		map[string]any{
			"id":         p.Id,
			"channel_id": p.ChannelId,
			"root_id":    p.RootId,
			"message":    p.Message,
			"create_at":  formatMilli(p.CreateAt),
		},
	)
}
