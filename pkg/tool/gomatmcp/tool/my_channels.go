package tool

import (
	"context"
	"fmt"
	"github.com/funtimecoding/go-library/pkg/generative/mark/response"
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mattermost/mattermost/server/public/model"
	"sort"
)

type myChannelsArguments struct {
	Limit      int    `json:"limit"`
	TypeFilter string `json:"type"`
	Since      string `json:"since"`
}

func (t *Tool) MyChannels(
	_ context.Context,
	_ mcp.CallToolRequest,
	arguments myChannelsArguments,
) (result *mcp.CallToolResult, e error) {
	defer func() {
		if r := recover(); r != nil {
			result = mcp.NewToolResultError(fmt.Sprintf("%v", r))
			e = nil
		}
	}()
	limit := arguments.Limit

	if limit <= 0 {
		limit = 30
	}

	me := t.client.Me()
	team := t.client.DefaultTeam()
	channels := t.client.Channels(team, me)
	sort.Slice(
		channels,
		func(a, b int) bool {
			return channels[a].LastPostAt > channels[b].LastPostAt
		},
	)

	if arguments.Since != "" {
		since, f := parseSince(arguments.Since)

		if f != nil {
			return response.Fail(
				"invalid since format: %v",
				f,
			)
		}

		sinceMilli := since.UnixMilli()
		var filtered []*model.Channel

		for _, c := range channels {
			if c.LastPostAt >= sinceMilli {
				filtered = append(filtered, c)
			}
		}

		channels = filtered
	}

	if arguments.TypeFilter != "" {
		var filtered []*model.Channel

		for _, c := range channels {
			if string(c.Type) == arguments.TypeFilter {
				filtered = append(filtered, c)
			}
		}

		channels = filtered
	}

	if len(channels) > limit {
		channels = channels[:limit]
	}

	type row struct {
		Identifier  string `json:"identifier"`
		Name        string `json:"name"`
		DisplayName string `json:"display_name"`
		Type        string `json:"type"`
		LastPostAt  string `json:"last_post_at"`
	}
	rows := make([]row, len(channels))

	for i, c := range channels {
		typeName := ""

		switch c.Type {
		case model.ChannelTypeOpen:
			typeName = "public"
		case model.ChannelTypePrivate:
			typeName = "private"
		case model.ChannelTypeDirect:
			typeName = "dm"
		case model.ChannelTypeGroup:
			typeName = "group_dm"
		}

		displayName := t.channelDisplayName(c)
		rows[i] = row{
			Identifier:  c.Id,
			Name:        c.Name,
			DisplayName: displayName,
			Type:        typeName,
			LastPostAt:  formatMilli(c.LastPostAt),
		}
	}

	return response.SuccessAny(rows)
}
