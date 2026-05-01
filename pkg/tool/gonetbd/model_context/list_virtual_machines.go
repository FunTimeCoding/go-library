package model_context

import (
	"context"
	"github.com/funtimecoding/go-library/pkg/generative/mark/response"
	"github.com/funtimecoding/go-library/pkg/tool/gonetbd/constant"
	"github.com/funtimecoding/go-library/pkg/tool/gonetbd/convert"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) listVirtualMachines(
	_ context.Context,
	_ mcp.CallToolRequest,
) (*mcp.CallToolResult, error) {
	result, e := s.client.VirtualMachines()

	if e != nil {
		return s.captureFail(e, constant.Unreachable)
	}

	return response.SuccessAny(convert.VirtualMachines(result))
}
