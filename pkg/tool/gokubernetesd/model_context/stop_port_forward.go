package model_context

import (
	"context"
	"github.com/funtimecoding/go-library/pkg/generative/mark/response"
	"github.com/funtimecoding/go-library/pkg/tool/gokubernetesd/model_context/argument"
	"github.com/funtimecoding/go-library/pkg/tool/gokubernetesd/service/port_forward_state"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) StopPortForward(
	_ context.Context,
	_ mcp.CallToolRequest,
	a argument.StopPortForward,
) (*mcp.CallToolResult, error) {
	if a.ID == "" {
		return response.Fail("id is required")
	}

	v, okay := s.service.StopPortForward(a.ID)

	if !okay {
		return response.Fail("port forward not found: %s", a.ID)
	}

	state := v.(*port_forward_state.PortForwardState)
	close(state.Stop)

	return response.Success("stopped port forward %s", a.ID)
}
