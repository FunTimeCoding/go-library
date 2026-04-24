package tool

import (
	"context"
	"fmt"
	"github.com/funtimecoding/go-library/pkg/generative/mark/response"
	"github.com/funtimecoding/go-library/pkg/tool/gomatmcp/argument"
	"github.com/mark3labs/mcp-go/mcp"
)

func (t *Tool) GetUserProfile(
	_ context.Context,
	_ mcp.CallToolRequest,
	a argument.GetUserProfile,
) (result *mcp.CallToolResult, e error) {
	defer func() {
		if r := recover(); r != nil {
			result = mcp.NewToolResultError(fmt.Sprintf("%v", r))
			e = nil
		}
	}()

	if a.UserID == "" {
		return response.Fail("user_id is required")
	}

	u := t.client.User(a.UserID)

	return response.SuccessAny(
		map[string]any{
			"id":         u.Id,
			"username":   u.Username,
			"first_name": u.FirstName,
			"last_name":  u.LastName,
			"nickname":   u.Nickname,
			"email":      u.Email,
			"position":   u.Position,
			"roles":      u.Roles,
			"is_bot":     u.IsBot,
			"create_at":  formatMilli(u.CreateAt),
		},
	)
}
