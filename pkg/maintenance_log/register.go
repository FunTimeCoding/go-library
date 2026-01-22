package maintenance_log

import "github.com/mark3labs/mcp-go/mcp"

func (s *Server) register() {
	s.server.AddTool(
		mcp.NewTool(
			"add_entry",
			mcp.WithDescription("Add a new maintenance log entry"),
			mcp.WithString(
				"action",
				mcp.Required(),
				mcp.Description(
					"Brief description of what was done (e.g., 'restarted service', 'deployed v1.2.3')",
				),
			),
			mcp.WithString(
				"user",
				mcp.Required(),
				mcp.Description(
					"Who performed the action (username, email, or identifier)",
				),
			),
			mcp.WithString(
				"system",
				mcp.Description(
					"System identifier (node name, hostname, IP address) - optional",
				),
			),
			mcp.WithString(
				"service",
				mcp.Description(
					"Service identifier (project name, deployment, application) - optional",
				),
			),
			mcp.WithString(
				"description",
				mcp.Description(
					"Additional details or context about the action - optional",
				),
			),
			mcp.WithString(
				"timestamp",
				mcp.Description(
					"When the event happened (RFC3339 format) - defaults to now if not provided",
				),
			),
		),
		s.add,
	)
	s.server.AddTool(
		mcp.NewTool(
			"list_entries",
			mcp.WithDescription(
				"List maintenance log entries with optional filters",
			),
			mcp.WithString(
				"system",
				mcp.Description("Filter by system identifier"),
			),
			mcp.WithString(
				"service",
				mcp.Description("Filter by service identifier"),
			),
			mcp.WithString(
				"user",
				mcp.Description("Filter by user"),
			),
			mcp.WithString(
				"since",
				mcp.Description(
					"Show entries since this time (RFC3339 format)",
				),
			),
			mcp.WithString(
				"until",
				mcp.Description(
					"Show entries until this time (RFC3339 format)",
				),
			),
			mcp.WithNumber(
				"limit",
				mcp.Description(
					"Maximum number of entries to return (default: all)",
				),
			),
		),
		s.list,
	)
}
