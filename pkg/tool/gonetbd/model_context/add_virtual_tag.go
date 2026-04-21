package model_context

import (
	"context"
	"github.com/funtimecoding/go-library/pkg/generative/mark/response"
	"github.com/funtimecoding/go-library/pkg/tool/gonetbd/constant"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) addVirtualTag(
	_ context.Context,
	r mcp.CallToolRequest,
) (*mcp.CallToolResult, error) {
	name, f := r.RequireString(constant.VirtualMachine)

	if f != nil {
		return response.Fail("virtual_machine is required: %v", f)
	}

	tag, g := r.RequireString(constant.Tag)

	if g != nil {
		return response.Fail("tag is required: %v", g)
	}

	return response.SuccessAny(s.client.AddVirtualTag(name, tag))
}
