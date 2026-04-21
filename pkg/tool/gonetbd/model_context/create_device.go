package model_context

import (
	"context"
	"github.com/funtimecoding/go-library/pkg/generative/mark/response"
	"github.com/funtimecoding/go-library/pkg/generative/model_context/parameter"
	"github.com/funtimecoding/go-library/pkg/netbox/tenant"
	"github.com/funtimecoding/go-library/pkg/tool/gonetbd/constant"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) createDevice(
	_ context.Context,
	r mcp.CallToolRequest,
) (*mcp.CallToolResult, error) {
	name, f := r.RequireString(parameter.Name)

	if f != nil {
		return response.Fail("name is required: %v", f)
	}

	roleName, g := r.RequireString(constant.Role)

	if g != nil {
		return response.Fail("role is required: %v", g)
	}

	typeName, h := r.RequireString(constant.Type)

	if h != nil {
		return response.Fail("type is required: %v", h)
	}

	siteName, i := r.RequireString(constant.Site)

	if i != nil {
		return response.Fail("site is required: %v", i)
	}

	role := s.client.DeviceRoleByName(roleName)
	deviceType := s.client.DeviceTypeByName(typeName)
	site := s.client.SiteByName(siteName)
	var ten *tenant.Tenant

	if t := r.GetString(constant.Tenant, ""); t != "" {
		ten = s.client.TenantByName(t)
	}

	return response.SuccessAny(
		s.client.CreateDevice(name, role, nil, deviceType, site, ten),
	)
}
