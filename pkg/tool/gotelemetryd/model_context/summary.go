package model_context

import (
	"context"
	"github.com/funtimecoding/go-library/pkg/generative/mark/response"
	"github.com/funtimecoding/go-library/pkg/tool/gotelemetryd/constant"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) summary(
	_ context.Context,
	r mcp.CallToolRequest,
) (*mcp.CallToolResult, error) {
	since := r.GetString(constant.Since, "")
	until := r.GetString(constant.Until, "")
	groupBy := r.GetString(constant.GroupBy, constant.Tool)
	rows := s.store.Summary(since, until, groupBy)
	type entry struct {
		Tool    string `json:"tool"`
		Surface string `json:"surface,omitempty"`
		Count   int64  `json:"count"`
	}
	entries := make([]entry, len(rows))

	for i, row := range rows {
		entries[i] = entry{
			Tool:    row.Tool,
			Surface: row.Surface,
			Count:   row.Count,
		}
	}

	return response.SuccessAny(entries)
}
