package model_context

import (
	"context"
	markResponse "github.com/funtimecoding/go-library/pkg/generative/mark/response"
	"github.com/funtimecoding/go-library/pkg/tool/gotelemetryd/constant"
	"github.com/funtimecoding/go-library/pkg/tool/gotelemetryd/model_context/response"
	"github.com/funtimecoding/go-library/pkg/tool/gotelemetryd/store"
	"github.com/mark3labs/mcp-go/mcp"
	"time"
)

func (s *Server) query(
	_ context.Context,
	r mcp.CallToolRequest,
) (*mcp.CallToolResult, error) {
	o := store.NewQueryOption()
	o.Tool = r.GetString(constant.Tool, "")
	o.Surface = r.GetString(constant.Surface, "")
	o.Actor = r.GetString(constant.Actor, "")
	o.Kind = r.GetString(constant.Kind, "")
	o.Since = r.GetString(constant.Since, "")
	o.Until = r.GetString(constant.Until, "")
	o.Limit = r.GetInt(constant.Limit, 50)
	events := s.store.Recent(o)
	entries := make([]response.Event, len(events))

	for i, e := range events {
		entries[i] = response.Event{
			Tool:      e.Tool,
			Surface:   e.Surface,
			Actor:     e.Actor,
			Outcome:   e.Outcome,
			Kind:      e.Kind,
			Duration:  e.DurationMillisecond,
			CreatedAt: e.CreatedAt.Format(time.RFC3339),
		}
	}

	return markResponse.SuccessAny(entries)
}
