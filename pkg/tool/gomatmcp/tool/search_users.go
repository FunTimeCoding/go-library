package tool

import (
	"context"
	"fmt"
	"github.com/funtimecoding/go-library/pkg/generative/mark/response"
	"github.com/funtimecoding/go-library/pkg/tool/gomatmcp/argument"
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mattermost/mattermost/server/public/model"
)

func (t *Tool) SearchUsers(
	_ context.Context,
	_ mcp.CallToolRequest,
	a argument.SearchUsers,
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

	team := t.client.DefaultTeam()
	autocomplete, _, f := t.client.Nested().AutocompleteUsersInTeam(
		t.client.Context(),
		team.Id,
		a.Query,
		20,
		model.HeaderEtagClient,
	)

	if f != nil {
		return response.Fail("search failed: %v", f)
	}

	type row struct {
		Identifier string `json:"identifier"`
		Username   string `json:"username"`
		FirstName  string `json:"first_name,omitempty"`
		LastName   string `json:"last_name,omitempty"`
		Nickname   string `json:"nickname,omitempty"`
		Email      string `json:"email,omitempty"`
	}
	var rows []row

	for _, u := range autocomplete.Users {
		rows = append(
			rows,
			row{
				Identifier: u.Id,
				Username:   u.Username,
				FirstName:  u.FirstName,
				LastName:   u.LastName,
				Nickname:   u.Nickname,
				Email:      u.Email,
			},
		)
	}

	return response.SuccessAny(rows)
}
