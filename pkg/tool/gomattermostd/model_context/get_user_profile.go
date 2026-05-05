package model_context

import (
	"context"
	"github.com/funtimecoding/go-library/pkg/generative/mark/response"
	"github.com/funtimecoding/go-library/pkg/tool/gomattermostd/constant"
	"github.com/funtimecoding/go-library/pkg/tool/gomattermostd/model_context/argument"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) GetUserProfile(
	_ context.Context,
	_ mcp.CallToolRequest,
	a argument.GetUserProfile,
) (*mcp.CallToolResult, error) {
	if a.UserID == "" {
		return response.Fail("user_id is required")
	}

	u, e := s.client.User(a.UserID)

	if e != nil {
		return s.captureFail(e, constant.Unreachable)
	}

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
