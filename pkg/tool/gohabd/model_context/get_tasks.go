package model_context

import (
	"context"
	"github.com/funtimecoding/go-library/pkg/generative/mark/response"
	"github.com/funtimecoding/go-library/pkg/tool/gohabd/constant"
	"github.com/funtimecoding/go-library/pkg/tool/gohabd/convert"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) getTasks(
	_ context.Context,
	r mcp.CallToolRequest,
) (*mcp.CallToolResult, error) {
	taskType := r.GetString(constant.TaskType, "")

	return response.SuccessAny(
		convert.Tasks(s.habitica.Tasks(taskType)),
	)
}
