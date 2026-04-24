package tool

import (
	"context"
	"fmt"
	"github.com/funtimecoding/go-library/pkg/generative/mark/response"
	"github.com/funtimecoding/go-library/pkg/tool/gomatmcp/argument"
	"github.com/mark3labs/mcp-go/mcp"
)

func (t *Tool) GetUsers(
	_ context.Context,
	_ mcp.CallToolRequest,
	a argument.GetUsers,
) (result *mcp.CallToolResult, e error) {
	defer func() {
		if r := recover(); r != nil {
			result = mcp.NewToolResultError(fmt.Sprintf("%v", r))
			e = nil
		}
	}()
	limit := a.Limit

	if limit <= 0 {
		limit = 100
	}

	page := a.Page

	if page < 0 {
		page = 0
	}

	users := t.client.Users(page, limit)
	type row struct {
		ID        string `json:"id"`
		Username  string `json:"username"`
		FirstName string `json:"first_name"`
		LastName  string `json:"last_name"`
		Nickname  string `json:"nickname"`
		Email     string `json:"email"`
		IsBot     bool   `json:"is_bot"`
	}
	rows := make([]row, len(users))

	for i, u := range users {
		rows[i] = row{
			ID:        u.Id,
			Username:  u.Username,
			FirstName: u.FirstName,
			LastName:  u.LastName,
			Nickname:  u.Nickname,
			Email:     u.Email,
			IsBot:     u.IsBot,
		}
	}

	return response.SuccessAny(
		map[string]any{
			"users":    rows,
			"page":     page,
			"per_page": limit,
		},
	)
}
