package model_context

import (
	"context"
	"github.com/funtimecoding/go-library/pkg/generative/mark/response"
	"github.com/funtimecoding/go-library/pkg/generative/model_context/parameter"
	"github.com/funtimecoding/go-library/pkg/tool/gonetbd/constant"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) createCluster(
	_ context.Context,
	r mcp.CallToolRequest,
) (*mcp.CallToolResult, error) {
	name, f := r.RequireString(parameter.Name)

	if f != nil {
		return response.Fail("name is required: %v", f)
	}

	typeName, g := r.RequireString(constant.Type)

	if g != nil {
		return response.Fail("type is required: %v", g)
	}

	siteName, h := r.RequireString(constant.Site)

	if h != nil {
		return response.Fail("site is required: %v", h)
	}

	t := s.client.ClusterTypeByName(typeName)
	site := s.client.SiteByName(siteName)

	return response.SuccessAny(s.client.CreateCluster(name, t, site))
}
