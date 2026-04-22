package tool

import (
	"context"
	"fmt"
	"github.com/funtimecoding/go-library/pkg/generative/mark/response"
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mattermost/mattermost/server/public/model"
)

type sendDirectMessageArguments struct {
	Username string `json:"username"`
	Message  string `json:"message"`
}

func (t *Tool) SendDirectMessage(
	_ context.Context,
	_ mcp.CallToolRequest,
	arguments sendDirectMessageArguments,
) (result *mcp.CallToolResult, e error) {
	defer func() {
		if r := recover(); r != nil {
			result = mcp.NewToolResultError(fmt.Sprintf("%v", r))
			e = nil
		}
	}()

	if arguments.Username == "" {
		return response.Fail("username is required")
	}

	if arguments.Message == "" {
		return response.Fail("message is required")
	}

	user, resp, f := t.client.Nested().GetUserByUsername(
		t.client.Context(),
		arguments.Username,
		model.HeaderEtagClient,
	)

	if f != nil {
		if resp != nil && resp.StatusCode == 404 {
			return response.Fail(
				"user not found: %s",
				arguments.Username,
			)
		}

		return response.Fail("user lookup failed: %v", f)
	}

	p := t.client.DirectMessage(user, arguments.Message)

	return response.SuccessAny(
		map[string]any{
			"id":         p.Id,
			"channel_id": p.ChannelId,
			"to":         arguments.Username,
			"message":    p.Message,
			"create_at":  formatMilli(p.CreateAt),
		},
	)
}
