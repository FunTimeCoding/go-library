package model_context

import (
	"context"
	"fmt"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) getStatus(
	_ context.Context,
	_ mcp.CallToolRequest,
) (*mcp.CallToolResult, error) {
	totalRecords := s.store.Count()
	lastPoll := s.poller.LastPoll()

	if lastPoll.IsZero() {
		return mcp.NewToolResultText(
			fmt.Sprintf("Total records: %d\nLast poll: never", totalRecords),
		), nil
	}

	return mcp.NewToolResultText(
		fmt.Sprintf(
			"Total records: %d\nLast poll: %s",
			totalRecords,
			lastPoll.Format("2006-01-02T15:04:05Z07:00"),
		),
	), nil
}
