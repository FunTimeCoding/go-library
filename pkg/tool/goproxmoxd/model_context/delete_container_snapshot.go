package model_context

import (
	"context"
	"github.com/funtimecoding/go-library/pkg/generative/mark/response"
	"github.com/funtimecoding/go-library/pkg/tool/goproxmoxd/model_context/argument"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) DeleteContainerSnapshot(
	_ context.Context,
	_ mcp.CallToolRequest,
	a argument.DeleteContainerSnapshot,
) (*mcp.CallToolResult, error) {
	if a.Identifier == 0 {
		return response.Fail("identifier is required")
	}

	if a.Name == "" {
		return response.Fail("name is required")
	}

	c, e := s.findContainer(a.Identifier, a.Node)

	if e != nil {
		return s.captureDetail(e)
	}

	if c == nil {
		return response.Fail("container %d not found", a.Identifier)
	}

	task, e := s.client.DeleteContainerSnapshot(c, a.Name)

	if e != nil {
		return s.captureDetail(e)
	}

	return response.SuccessAny(map[string]string{"task_id": string(task.UPID)})
}
