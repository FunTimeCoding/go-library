package model_context

import (
	"context"
	"github.com/funtimecoding/go-library/pkg/generative/mark/response"
	"github.com/funtimecoding/go-library/pkg/tool/goansibled/constant"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) triggerRun(
	_ context.Context,
	r mcp.CallToolRequest,
) (*mcp.CallToolResult, error) {
	playbook := r.GetString(constant.Playbook, "")
	var playbooks []string

	if playbook != "" {
		playbooks = []string{playbook}
	}

	e := s.runner.Trigger(playbooks)

	if e != nil {
		return response.Fail(e.Error())
	}

	return response.Success("run queued")
}
