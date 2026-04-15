package model_context

import "github.com/mark3labs/mcp-go/mcp"

func (s *Server) register() {
	s.server.AddTool(
		mcp.NewTool(
			"netbox_list_devices",
			mcp.WithDescription("List NetBox devices, optionally filtered by name"),
			mcp.WithString(
				"query",
				mcp.Description("Filter by name (case-insensitive contains)"),
			),
		),
		s.listDevices,
	)
	s.server.AddTool(
		mcp.NewTool(
			"netbox_get_device",
			mcp.WithDescription("Get a NetBox device by exact name"),
			mcp.WithString("name", mcp.Required(), mcp.Description("Device name")),
		),
		s.getDevice,
	)
	s.server.AddTool(
		mcp.NewTool(
			"netbox_list_interfaces",
			mcp.WithDescription("List interfaces for a NetBox device"),
			mcp.WithString("name", mcp.Required(), mcp.Description("Device name")),
		),
		s.listInterfaces,
	)
	s.server.AddTool(
		mcp.NewTool(
			"netbox_list_addresses",
			mcp.WithDescription("List IP addresses assigned to a NetBox device"),
			mcp.WithString("name", mcp.Required(), mcp.Description("Device name")),
		),
		s.listAddresses,
	)
	s.server.AddTool(
		mcp.NewTool(
			"netbox_list_sites",
			mcp.WithDescription("List all NetBox sites"),
		),
		s.listSites,
	)
}
