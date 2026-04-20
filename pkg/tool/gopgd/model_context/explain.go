package model_context

import (
	"context"
	"fmt"
	"github.com/funtimecoding/go-library/pkg/generative/mark/response"
	"github.com/mark3labs/mcp-go/mcp"
)

type explainArguments struct {
	SQL     string `json:"sql"`
	Analyze bool   `json:"analyze"`
}

func (s *Server) explain(
	x context.Context,
	_ mcp.CallToolRequest,
	a explainArguments,
) (*mcp.CallToolResult, error) {
	instance, ok := s.activeInstance(x)

	if !ok {
		return response.Fail(
			"no instance selected — use use_instance first",
		)
	}

	if a.SQL == "" {
		return response.Fail("sql is required")
	}

	prefix := "EXPLAIN"

	if a.Analyze {
		prefix = "EXPLAIN ANALYZE"
	}

	v, e := s.store.Query(x, instance, fmt.Sprintf("%s %s", prefix, a.SQL))

	if e != nil {
		return response.Fail("explain: %v", e)
	}

	return response.SuccessAny(v)
}
