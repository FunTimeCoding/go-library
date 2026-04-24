package tool

import (
	"context"
	"fmt"
	"github.com/funtimecoding/go-library/pkg/generative/mark/response"
	"github.com/funtimecoding/go-library/pkg/tool/gomatmcp/argument"
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mattermost/mattermost/server/public/model"
	"strings"
)

func (t *Tool) SearchChannels(
	_ context.Context,
	_ mcp.CallToolRequest,
	a argument.SearchChannels,
) (result *mcp.CallToolResult, e error) {
	defer func() {
		if r := recover(); r != nil {
			result = mcp.NewToolResultError(fmt.Sprintf("%v", r))
			e = nil
		}
	}()

	if a.Query == "" {
		return response.Fail("query is required")
	}

	me := t.client.Me()
	team := t.client.DefaultTeam()
	channels, _, f := t.client.Nested().AutocompleteChannelsForTeamForSearch(
		t.client.Context(),
		team.Id,
		a.Query,
	)

	if f != nil {
		return response.Fail("search failed: %v", f)
	}

	type row struct {
		Identifier  string `json:"identifier"`
		Name        string `json:"name"`
		DisplayName string `json:"display_name"`
		Type        string `json:"type"`
		LastPostAt  string `json:"last_post_at,omitempty"`
	}
	var rows []row

	for _, c := range channels {
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

		displayName := c.DisplayName

		if c.Type == model.ChannelTypeDirect && displayName == "" {
			parts := strings.Split(c.Name, "__")

			for _, p := range parts {
				if p != me.Id {
					other := t.client.User(p)
					displayName = other.Username

					break
				}
			}
		}

		r := row{
			Identifier:  c.Id,
			Name:        c.Name,
			DisplayName: displayName,
			Type:        typeName,
		}

		if c.LastPostAt > 0 {
			r.LastPostAt = formatMilli(c.LastPostAt)
		}

		rows = append(rows, r)
	}

	return response.SuccessAny(rows)
}
