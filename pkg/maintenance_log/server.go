package maintenance_log

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
	"time"
)

type Server struct {
	mcp   *server.MCPServer
	store *Store
}

func NewServer() (*Server, error) {
	store, err := NewStore()
	if err != nil {
		return nil, fmt.Errorf("failed to create store: %w", err)
	}

	mcpServer := server.NewMCPServer(
		"maintenance-log",
		"1.0.0",
		server.WithToolCapabilities(true),
	)

	s := &Server{
		mcp:   mcpServer,
		store: store,
	}

	s.registerTools()

	return s, nil
}

func (s *Server) registerTools() {
	s.mcp.AddTool(
		mcp.NewTool(
			"add_entry",
			mcp.WithDescription("Add a new maintenance log entry"),
			mcp.WithString(
				"action",
				mcp.Required(),
				mcp.Description("Brief description of what was done (e.g., 'restarted service', 'deployed v1.2.3')"),
			),
			mcp.WithString(
				"user",
				mcp.Required(),
				mcp.Description("Who performed the action (username, email, or identifier)"),
			),
			mcp.WithString(
				"system",
				mcp.Description("System identifier (node name, hostname, IP address) - optional"),
			),
			mcp.WithString(
				"service",
				mcp.Description("Service identifier (project name, deployment, application) - optional"),
			),
			mcp.WithString(
				"description",
				mcp.Description("Additional details or context about the action - optional"),
			),
			mcp.WithString(
				"timestamp",
				mcp.Description("When the event happened (RFC3339 format) - defaults to now if not provided"),
			),
		),
		s.handleAddEntry,
	)

	s.mcp.AddTool(
		mcp.NewTool(
			"list_entries",
			mcp.WithDescription("List maintenance log entries with optional filters"),
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
				mcp.Description("Show entries since this time (RFC3339 format)"),
			),
			mcp.WithString(
				"until",
				mcp.Description("Show entries until this time (RFC3339 format)"),
			),
			mcp.WithNumber(
				"limit",
				mcp.Description("Maximum number of entries to return (default: all)"),
			),
		),
		s.handleListEntries,
	)
}

func (s *Server) handleAddEntry(
	ctx context.Context,
	req mcp.CallToolRequest,
) (*mcp.CallToolResult, error) {
	action, err := req.RequireString("action")
	if err != nil {
		return mcp.NewToolResultError(
			fmt.Sprintf(
				"action is required: %v",
				err,
			),
		), nil
	}

	user, err := req.RequireString("user")
	if err != nil {
		return mcp.NewToolResultError(
			fmt.Sprintf(
				"user is required: %v",
				err,
			),
		), nil
	}

	entry := &Entry{
		Action: action,
		User:   user,
	}

	if system := req.GetString("system", ""); system != "" {
		entry.System = system
	}

	if service := req.GetString("service", ""); service != "" {
		entry.Service = service
	}

	if description := req.GetString("description", ""); description != "" {
		entry.Description = description
	}

	if timestampStr := req.GetString("timestamp", ""); timestampStr != "" {
		timestamp, parseErr := time.Parse(time.RFC3339, timestampStr)
		if parseErr != nil {
			return mcp.NewToolResultError(
				fmt.Sprintf(
					"invalid timestamp format: %v",
					parseErr,
				),
			), nil
		}
		entry.Timestamp = timestamp
	}

	if err := s.store.AddEntry(entry); err != nil {
		return mcp.NewToolResultError(
			fmt.Sprintf(
				"failed to add entry: %v",
				err,
			),
		), nil
	}

	data, _ := json.MarshalIndent(entry, "", "  ")
	return mcp.NewToolResultText(
		fmt.Sprintf(
			"Entry added successfully:\n%s",
			string(data),
		),
	), nil
}

func (s *Server) handleListEntries(
	ctx context.Context,
	req mcp.CallToolRequest,
) (*mcp.CallToolResult, error) {
	filter := &EntryFilter{}

	if system := req.GetString("system", ""); system != "" {
		filter.System = system
	}

	if service := req.GetString("service", ""); service != "" {
		filter.Service = service
	}

	if user := req.GetString("user", ""); user != "" {
		filter.User = user
	}

	if sinceStr := req.GetString("since", ""); sinceStr != "" {
		since, parseErr := time.Parse(time.RFC3339, sinceStr)
		if parseErr != nil {
			return mcp.NewToolResultError(
				fmt.Sprintf(
					"invalid since format: %v",
					parseErr,
				),
			), nil
		}
		filter.Since = since
	}

	if untilStr := req.GetString("until", ""); untilStr != "" {
		until, parseErr := time.Parse(time.RFC3339, untilStr)
		if parseErr != nil {
			return mcp.NewToolResultError(
				fmt.Sprintf(
					"invalid until format: %v",
					parseErr,
				),
			), nil
		}
		filter.Until = until
	}

	if limit := req.GetFloat("limit", 0); limit > 0 {
		filter.Limit = int(limit)
	}

	entries, err := s.store.ListEntries(filter)
	if err != nil {
		return mcp.NewToolResultError(
			fmt.Sprintf(
				"failed to list entries: %v",
				err,
			),
		), nil
	}

	data, _ := json.MarshalIndent(entries, "", "  ")
	return mcp.NewToolResultText(
		fmt.Sprintf(
			"Found %d entries:\n%s",
			len(entries),
			string(data),
		),
	), nil
}

func (s *Server) MCPServer() *server.MCPServer {
	return s.mcp
}
