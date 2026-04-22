package tool

import (
	"context"
	"fmt"
	"github.com/funtimecoding/go-library/pkg/generative/mark/response"
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mattermost/mattermost/server/public/model"
)

type postMessageArguments struct {
	ChannelID   string `json:"channel_id"`
	ChannelName string `json:"channel_name"`
	Message     string `json:"message"`
}

func (t *Tool) PostMessage(
	_ context.Context,
	_ mcp.CallToolRequest,
	a postMessageArguments,
) (result *mcp.CallToolResult, e error) {
	defer func() {
		if r := recover(); r != nil {
			result = mcp.NewToolResultError(fmt.Sprintf("%v", r))
			e = nil
		}
	}()

	if a.Message == "" {
		return response.Fail("message is required")
	}

	if a.ChannelID == "" && a.ChannelName == "" {
		return response.Fail("channel_id or channel_name is required")
	}

	var ch *model.Channel

	if a.ChannelName != "" {
		ch = t.client.TeamChannel(a.ChannelName)
	} else {
		ch = t.client.Channel(a.ChannelID)
	}

	p := t.client.PostSimple(ch, a.Message)

	return response.SuccessAny(
		map[string]any{
			"id":         p.Id,
			"channel_id": p.ChannelId,
			"message":    p.Message,
			"create_at":  formatMilli(p.CreateAt),
		},
	)
}
