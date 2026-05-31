package model_context

import (
	"context"
	"github.com/funtimecoding/go-library/pkg/generative/mark/response"
	"github.com/funtimecoding/go-library/pkg/tool/goproxmoxd/model_context/argument"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) ListContainerSnapshots(
	_ context.Context,
	_ mcp.CallToolRequest,
	a argument.ListContainerSnapshots,
) (*mcp.CallToolResult, error) {
	if a.Identifier == 0 {
		return response.Fail("identifier is required")
	}

	c, e := s.findContainer(a.Identifier, a.Node)

	if e != nil {
		return s.captureDetail(e)
	}

	if c == nil {
		return response.Fail("container %d not found", a.Identifier)
	}

	snapshots, e := s.client.ContainerSnapshots(c)

	if e != nil {
		return s.captureDetail(e)
	}

	return response.SuccessAny(snapshots)
}
