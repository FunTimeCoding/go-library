package tool

import (
	"context"
	"fmt"
	"github.com/funtimecoding/go-library/pkg/chat/mattermost/post"
	"github.com/funtimecoding/go-library/pkg/generative/mark/response"
	"github.com/funtimecoding/go-library/pkg/tool/gomatmcp/argument"
	"github.com/mark3labs/mcp-go/mcp"
)

func (t *Tool) SearchMessages(
	_ context.Context,
	_ mcp.CallToolRequest,
	a argument.SearchMessages,
) (result *mcp.CallToolResult, e error) {
	defer func() {
		if r := recover(); r != nil {
			result = mcp.NewToolResultError(fmt.Sprintf("%v", r))
			e = nil
		}
	}()

	if a.Terms == "" {
		return response.Fail("terms is required")
	}

	team := t.client.DefaultTeam()
	list, _, f := t.client.Nested().SearchPosts(
		t.client.Context(),
		team.Id,
		a.Terms,
		false,
	)

	if f != nil {
		return response.Fail("search failed: %v", f)
	}

	posts := post.NewSlice(post.FromList(list, true))
	t.client.Enrich(posts)
	type row struct {
		Identifier string   `json:"identifier"`
		Channel    string   `json:"channel"`
		Username   string   `json:"username"`
		Message    string   `json:"message"`
		CreateAt   string   `json:"create_at"`
		FileIds    []string `json:"file_ids,omitempty"`
	}
	var rows []row

	for _, p := range posts {
		r := row{
			Identifier: p.Raw.Id,
			Message:    p.Message,
			CreateAt:   formatTime(p.Create),
		}

		if p.User != nil {
			r.Username = p.User.Username
		}

		if p.Raw.ChannelId != "" {
			ch := t.client.Channel(p.Raw.ChannelId)
			r.Channel = t.channelDisplayName(ch)
		}

		if len(p.Raw.FileIds) > 0 {
			r.FileIds = p.Raw.FileIds
		}

		rows = append(rows, r)
	}

	return response.SuccessAny(
		map[string]any{
			"results": rows,
			"count":   len(rows),
		},
	)
}
