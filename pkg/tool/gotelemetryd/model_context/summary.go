package model_context

import (
	"context"
	markResponse "github.com/funtimecoding/go-library/pkg/generative/mark/response"
	"github.com/funtimecoding/go-library/pkg/tool/gotelemetryd/constant"
	"github.com/funtimecoding/go-library/pkg/tool/gotelemetryd/model_context/response"
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
	entries := make([]response.Summary, len(rows))

	for i, row := range rows {
		entries[i] = response.Summary{
			Tool:    row.Tool,
			Surface: row.Surface,
			Kind:    row.Kind,
			Count:   row.Count,
		}
	}

	return markResponse.SuccessAny(entries)
}
