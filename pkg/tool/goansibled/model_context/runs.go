package model_context

import (
	"context"
	mark "github.com/funtimecoding/go-library/pkg/generative/mark/response"
	"github.com/funtimecoding/go-library/pkg/time"
	"github.com/funtimecoding/go-library/pkg/tool/goansibled/constant"
	"github.com/funtimecoding/go-library/pkg/tool/goansibled/model_context/response"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) runs(
	_ context.Context,
	r mcp.CallToolRequest,
) (*mcp.CallToolResult, error) {
	limit := r.GetInt(constant.Limit, 20)
	result, e := s.store.Recent(limit)

	if e != nil {
		return s.captureFail(e, constant.RecentRunsFailed)
	}

	summaries := make([]response.RunSummary, len(result))

	for i, v := range result {
		summaries[i] = response.RunSummary{
			ID:                  v.ID,
			CreatedAt:           v.CreatedAt.Format(time.DateSecond),
			Playbook:            v.Playbook,
			TriggerSource:       v.TriggerSource,
			DurationMillisecond: v.DurationMillisecond,
			Status:              v.Status,
			GitHead:             v.GitHead,
		}
	}

	return mark.SuccessAny(summaries)
}
