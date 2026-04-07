package tool

import (
	"context"
	"fmt"
	"github.com/funtimecoding/go-library/pkg/generative/mark/response"
	"github.com/mark3labs/mcp-go/mcp"
)

type addReactionArguments struct {
	PostID    string `json:"post_id"`
	EmojiName string `json:"emoji_name"`
}

func (t *Tool) AddReaction(
	_ context.Context,
	_ mcp.CallToolRequest,
	arguments addReactionArguments,
) (result *mcp.CallToolResult, e error) {
	defer func() {
		if r := recover(); r != nil {
			result = mcp.NewToolResultError(fmt.Sprintf("%v", r))
			e = nil
		}
	}()

	if arguments.PostID == "" {
		return response.Fail("post_id is required")
	}

	if arguments.EmojiName == "" {
		return response.Fail("emoji_name is required")
	}

	p := t.client.FindPost(arguments.PostID)
	t.client.React(p, arguments.EmojiName)

	return response.SuccessAny(
		map[string]any{
			"post_id":    arguments.PostID,
			"emoji_name": arguments.EmojiName,
		},
	)
}
