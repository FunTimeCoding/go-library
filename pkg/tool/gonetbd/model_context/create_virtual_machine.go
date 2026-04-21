package model_context

import (
	"context"
	"github.com/funtimecoding/go-library/pkg/generative/mark/response"
	"github.com/funtimecoding/go-library/pkg/generative/model_context/parameter"
	"github.com/funtimecoding/go-library/pkg/tool/gonetbd/constant"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) createVirtualMachine(
	_ context.Context,
	r mcp.CallToolRequest,
) (*mcp.CallToolResult, error) {
	name, f := r.RequireString(parameter.Name)

	if f != nil {
		return response.Fail("name is required: %v", f)
	}

	clusterName, g := r.RequireString(constant.Cluster)

	if g != nil {
		return response.Fail("cluster is required: %v", g)
	}

	cl := s.client.ClusterByName(clusterName)

	return response.SuccessAny(s.client.CreateVirtualMachine(name, cl))
}
