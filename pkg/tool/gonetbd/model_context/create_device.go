package model_context

import (
	"context"
	"fmt"
	"github.com/funtimecoding/go-library/pkg/netbox/tenant"
	"github.com/funtimecoding/go-library/pkg/notation"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) createDevice(
	_ context.Context,
	r mcp.CallToolRequest,
) (*mcp.CallToolResult, error) {
	name, f := r.RequireString("name")

	if f != nil {
		return mcp.NewToolResultError(
			fmt.Sprintf(
				"name is required: %v",
				f,
			),
		), nil
	}

	roleName, f := r.RequireString("role")

	if f != nil {
		return mcp.NewToolResultError(
			fmt.Sprintf(
				"role is required: %v",
				f,
			),
		), nil
	}

	typeName, f := r.RequireString("type")

	if f != nil {
		return mcp.NewToolResultError(
			fmt.Sprintf(
				"type is required: %v",
				f,
			),
		), nil
	}

	siteName, f := r.RequireString("site")

	if f != nil {
		return mcp.NewToolResultError(
			fmt.Sprintf(
				"site is required: %v",
				f,
			),
		), nil
	}

	role := s.client.DeviceRoleByName(roleName)
	deviceType := s.client.DeviceTypeByName(typeName)
	site := s.client.SiteByName(siteName)
	var ten *tenant.Tenant

	if t := r.GetString("tenant", ""); t != "" {
		ten = s.client.TenantByName(t)
	}

	return mcp.NewToolResultText(
		notation.MarshalIndent(
			s.client.CreateDevice(name, role, nil, deviceType, site, ten),
		),
	), nil
}
