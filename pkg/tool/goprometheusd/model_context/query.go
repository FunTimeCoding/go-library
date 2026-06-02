package model_context

import (
	"context"
	"github.com/funtimecoding/go-library/pkg/generative/mark/response"
	"github.com/funtimecoding/go-library/pkg/tool/goprometheusd/convert"
	"github.com/funtimecoding/go-library/pkg/tool/goprometheusd/model_context/argument"
	"github.com/mark3labs/mcp-go/mcp"
	"time"
)

func (s *Server) query(
	x context.Context,
	_ mcp.CallToolRequest,
	a argument.Query,
) (*mcp.CallToolResult, error) {
	instance, okay := s.activeInstance(x)

	if !okay {
		return response.Fail(
			"no instance selected - use use_instance first",
		)
	}

	if a.Query == "" {
		return response.Fail("query is required")
	}

	v, e := s.service.Query(instance, a.Query, time.Now())

	if e != nil {
		return s.captureDetail(e)
	}

	return response.SuccessAny(convert.QueryResult(v))
}
