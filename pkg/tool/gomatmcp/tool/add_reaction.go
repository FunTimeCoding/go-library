package tool

import (
	"context"
	"fmt"
	"github.com/funtimecoding/go-library/pkg/generative/mark/response"
	"github.com/funtimecoding/go-library/pkg/tool/gomatmcp/argument"
	"github.com/mark3labs/mcp-go/mcp"
)

func (t *Tool) AddReaction(
	_ context.Context,
	_ mcp.CallToolRequest,
	a argument.AddReaction,
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

	if a.EmojiName == "" {
		return response.Fail("emoji_name is required")
	}

	p := t.client.FindPost(a.PostID)
	t.client.React(p, a.EmojiName)

	return response.SuccessAny(
		map[string]any{
			"post_id":    a.PostID,
			"emoji_name": a.EmojiName,
		},
	)
}
