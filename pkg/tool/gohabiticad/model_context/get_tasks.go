package model_context

import (
	"context"
	"github.com/funtimecoding/go-library/pkg/generative/mark/response"
	"github.com/funtimecoding/go-library/pkg/tool/gohabiticad/constant"
	"github.com/funtimecoding/go-library/pkg/tool/gohabiticad/convert"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) getTasks(
	_ context.Context,
	r mcp.CallToolRequest,
) (*mcp.CallToolResult, error) {
	taskType := r.GetString(constant.TaskType, "")
	result, e := s.habitica.Tasks(taskType)

	if e != nil {
		return s.captureFail(e, constant.Unreachable)
	}

	return response.SuccessAny(convert.Tasks(result))
}
