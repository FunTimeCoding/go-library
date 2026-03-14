package model_context

import "github.com/mark3labs/mcp-go/mcp"

func (s *Server) register() {
	s.server.AddTool(
		mcp.NewTool(
			"get_alerts",
			mcp.WithDescription("Get alert log entries matching the given name"),
			mcp.WithString(
				"name",
				mcp.Required(),
				mcp.Description("Alert name to filter by"),
			),
		),
		s.getAlerts,
	)
	s.server.AddTool(
		mcp.NewTool(
			"get_recent_alerts",
			mcp.WithDescription(
				"Get alert log entries within a time range (defaults to last 1 hour)",
			),
			mcp.WithString(
				"start",
				mcp.Description(
					"Start of time range (RFC3339 format) - defaults to 1 hour ago",
				),
			),
			mcp.WithString(
				"end",
				mcp.Description(
					"End of time range (RFC3339 format) - defaults to now",
				),
			),
		),
		s.getRecentAlerts,
	)
	s.server.AddTool(
		mcp.NewTool(
			"get_status",
			mcp.WithDescription(
				"Get alert log service status including total records and last poll time",
			),
		),
		s.getStatus,
	)
}
