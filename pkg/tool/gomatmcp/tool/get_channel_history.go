package tool

import (
	"context"
	"fmt"
	"github.com/funtimecoding/go-library/pkg/chat/mattermost/post"
	"github.com/funtimecoding/go-library/pkg/generative/mark/response"
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mattermost/mattermost/server/public/model"
	"time"
)

type getChannelHistoryArguments struct {
	ChannelID   string `json:"channel_id"`
	ChannelName string `json:"channel_name"`
	Limit       int    `json:"limit"`
	Since       string `json:"since"`
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

	var posts []*post.Post

	if arguments.Since != "" {
		since, f := parseSince(arguments.Since)

		if f != nil {
			return response.Fail(
				"invalid since format: %v",
				f,
			)
		}

		posts = t.client.RecentPosts(ch, since.UnixMilli())
	} else {
		posts = t.client.PostsBefore(ch, time.Now(), limit)
		t.client.Enrich(posts)
	}

	type row struct {
		ID       string   `json:"id"`
		Username string   `json:"username"`
		Message  string   `json:"message"`
		CreateAt string   `json:"create_at"`
		RootID   *string  `json:"root_id,omitempty"`
		FileIds  []string `json:"file_ids,omitempty"`
	}
	rows := make([]row, len(posts))

	for i, p := range posts {
		r := row{
			ID:       p.Raw.Id,
			Message:  p.Message,
			CreateAt: formatTime(p.Create),
		}

		if p.User != nil {
			r.Username = p.User.Username
		}

		if p.Raw.RootId != "" {
			r.RootID = new(p.Raw.RootId)
		}

		if len(p.Raw.FileIds) > 0 {
			r.FileIds = p.Raw.FileIds
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
