package model_context

import (
	"context"
	"github.com/funtimecoding/go-library/pkg/generative/mark/response"
	"github.com/funtimecoding/go-library/pkg/habitica"
	"github.com/funtimecoding/go-library/pkg/tool/gohabd/constant"
	"github.com/funtimecoding/go-library/pkg/tool/gohabd/convert"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) createTask(
	_ context.Context,
	r mcp.CallToolRequest,
) (*mcp.CallToolResult, error) {
	taskType, f := r.RequireString(constant.TaskType)

	if f != nil {
		return response.Fail("task_type is required: %v", f)
	}

	text, g := r.RequireString(constant.Text)

	if g != nil {
		return response.Fail("text is required: %v", g)
	}

	body := habitica.CreateTaskBody{
		Type:  taskType,
		Text:  text,
		Notes: r.GetString(constant.Notes, ""),
	}

	return response.SuccessAny(
		convert.Task(s.habitica.CreateTask(body)),
	)
}
