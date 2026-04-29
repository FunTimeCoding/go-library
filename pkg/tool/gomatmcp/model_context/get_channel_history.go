package model_context

import (
	"context"
	"github.com/funtimecoding/go-library/pkg/chat/mattermost/post"
	"github.com/funtimecoding/go-library/pkg/generative/mark/response"
	"github.com/funtimecoding/go-library/pkg/tool/gomatmcp/argument"
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mattermost/mattermost/server/public/model"
	"time"
)

func (s *Server) GetChannelHistory(
	_ context.Context,
	_ mcp.CallToolRequest,
	a argument.GetChannelHistory,
) (*mcp.CallToolResult, error) {
	if a.ChannelID == "" && a.ChannelName == "" {
		return response.Fail("channel_id or channel_name is required")
	}

	limit := a.Limit

	if limit <= 0 {
		limit = 30
	}

	var ch *model.Channel

	if a.ChannelName != "" {
		var e error
		ch, e = s.client.TeamChannel(a.ChannelName)

		if e != nil {
			return s.captureFail(e, "channel not found")
		}
	} else {
		var e error
		ch, e = s.client.Channel(a.ChannelID)

		if e != nil {
			return s.captureFail(e, "channel not found")
		}
	}

	var posts []*post.Post
	var g error

	if a.Since != "" {
		since, f := parseSince(a.Since)

		if f != nil {
			return response.Fail(
				"invalid since format: %v",
				f,
			)
		}

		posts, g = s.client.RecentPosts(ch, since.UnixMilli())

		if g != nil {
			return s.captureFail(g, "get recent posts failed")
		}
	} else {
		posts, g = s.client.PostsBefore(ch, time.Now(), limit)

		if g != nil {
			return s.captureFail(g, "get posts failed")
		}

		h := s.client.Enrich(posts)

		if h != nil {
			return s.captureFail(h, "enrich posts failed")
		}
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
