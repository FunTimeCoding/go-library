package model_context

import (
	"context"
	"github.com/funtimecoding/go-library/pkg/generative/mark/response"
	"github.com/funtimecoding/go-library/pkg/tool/gopgd/argument"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) tableSizes(
	x context.Context,
	_ mcp.CallToolRequest,
	a argument.TableSizes,
) (*mcp.CallToolResult, error) {
	instance, okay := s.activeInstance(x)

	if !okay {
		return response.Fail(
			"no instance selected — use use_instance first",
		)
	}

	schema := a.Schema

	if schema == "" {
		schema = "public"
	}

	v, e := s.store.TableSizes(x, instance, schema)

	if e != nil {
		return response.Fail("table sizes: %v", e)
	}

	return response.SuccessAny(v)
}
