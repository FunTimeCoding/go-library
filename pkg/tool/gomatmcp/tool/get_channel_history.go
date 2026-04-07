package tool

import (
	"context"
	"fmt"
	"github.com/funtimecoding/go-library/pkg/generative/mark/response"
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mattermost/mattermost/server/public/model"
	"time"
)

type getChannelHistoryArguments struct {
	ChannelID   string `json:"channel_id"`
	ChannelName string `json:"channel_name"`
	Limit       int    `json:"limit"`
}

func (t *Tool) GetChannelHistory(
	_ context.Context,
	_ mcp.CallToolRequest,
	arguments getChannelHistoryArguments,
) (result *mcp.CallToolResult, e error) {
	defer func() {
		if r := recover(); r != nil {
			result = mcp.NewToolResultError(fmt.Sprintf("%v", r))
			e = nil
		}
	}()

	if arguments.ChannelID == "" && arguments.ChannelName == "" {
		return response.Fail("channel_id or channel_name is required")
	}

	limit := arguments.Limit

	if limit <= 0 {
		limit = 30
	}

	var ch *model.Channel

	if arguments.ChannelName != "" {
		ch = t.client.TeamChannel(arguments.ChannelName)
	} else {
		ch = t.client.Channel(arguments.ChannelID)
	}

	posts := t.client.PostsBefore(ch, time.Now(), limit)
	type row struct {
		ID       string  `json:"id"`
		UserID   string  `json:"user_id"`
		Message  string  `json:"message"`
		CreateAt string  `json:"create_at"`
		RootID   *string `json:"root_id,omitempty"`
	}
	rows := make([]row, len(posts))

	for i, p := range posts {
		r := row{
			ID:       p.Raw.Id,
			UserID:   p.Raw.UserId,
			Message:  p.Message,
			CreateAt: p.Create.Format(time.RFC3339),
		}

		if p.Raw.RootId != "" {
			r.RootID = new(p.Raw.RootId)
		}

		rows[i] = r
	}

	return response.SuccessAny(
		map[string]any{
			"posts":    rows,
			"per_page": limit,
		},
	)
}
