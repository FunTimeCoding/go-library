package model_context

import (
	"context"
	"github.com/funtimecoding/go-library/pkg/generative/mark/response"
	"github.com/funtimecoding/go-library/pkg/tool/goproxmoxd/model_context/argument"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) ListContainerSnapshots(
	x context.Context,
	_ mcp.CallToolRequest,
	a argument.ListContainerSnapshots,
) (*mcp.CallToolResult, error) {
	if a.Identifier == 0 {
		return response.Fail("identifier is required")
	}

	instance, e := s.service.ResolveInstance(s.activeInstanceName(x))

	if e != nil {
		return response.Fail("%s", e)
	}

	p, e := s.service.Client(instance)

	if e != nil {
		return s.captureDetail(e)
	}

	ct, e := findContainer(p, a.Identifier, a.Node)

	if e != nil {
		return s.captureDetail(e)
	}

	if ct == nil {
		return response.Fail("container %d not found", a.Identifier)
	}

	snapshots, e := p.ContainerSnapshots(ct)

	if e != nil {
		return s.captureDetail(e)
	}

	return response.SuccessAny(snapshots)
}
