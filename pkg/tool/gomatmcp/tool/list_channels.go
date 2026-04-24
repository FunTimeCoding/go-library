package tool

import (
	"context"
	"fmt"
	"github.com/funtimecoding/go-library/pkg/generative/mark/response"
	"github.com/funtimecoding/go-library/pkg/tool/gomatmcp/argument"
	"github.com/mark3labs/mcp-go/mcp"
)

func (t *Tool) ListChannels(
	_ context.Context,
	_ mcp.CallToolRequest,
	a argument.ListChannels,
) (result *mcp.CallToolResult, e error) {
	defer func() {
		if r := recover(); r != nil {
			result = mcp.NewToolResultError(fmt.Sprintf("%v", r))
			e = nil
		}
	}()
	limit := a.Limit

	if limit <= 0 {
		limit = 100
	}

	all := t.client.AllChannels()
	start := a.Page * limit

	if start >= len(all) {
		start = len(all)
	}

	end := start + limit

	if end > len(all) {
		end = len(all)
	}

	page := all[start:end]
	type row struct {
		ID          string `json:"id"`
		Name        string `json:"name"`
		DisplayName string `json:"display_name"`
		Type        string `json:"type"`
		Purpose     string `json:"purpose"`
		Header      string `json:"header"`
	}
	rows := make([]row, len(page))

	for i, c := range page {
		rows[i] = row{
			ID:          c.Id,
			Name:        c.Name,
			DisplayName: c.DisplayName,
			Type:        string(c.Type),
			Purpose:     c.Purpose,
			Header:      c.Header,
		}
	}

	return response.SuccessAny(
		map[string]any{
			"channels":    rows,
			"total_count": len(all),
			"page":        a.Page,
			"per_page":    limit,
		},
	)
}
