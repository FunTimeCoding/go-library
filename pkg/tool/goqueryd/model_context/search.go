package model_context

import (
	"context"
	"github.com/funtimecoding/go-library/pkg/generative/mark/response"
	"github.com/funtimecoding/go-library/pkg/generative/model_context/parameter"
	"github.com/funtimecoding/go-library/pkg/notation"
	"github.com/funtimecoding/go-library/pkg/tool/goqueryd/constant"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) search(
	_ context.Context,
	q mcp.CallToolRequest,
) (*mcp.CallToolResult, error) {
	query, e := q.RequireString(parameter.Query)

	if e != nil {
		return response.Fail("query is required: %v", e)
	}

	metadata := extractMetadata(q)
	sourceType := q.GetString(constant.SourceType, "")

	if sourceType != "" {
		if metadata == nil {
			metadata = map[string]string{}
		}

		metadata[constant.SourceType] = sourceType
	}

	outcome := s.service.Search(
		query,
		int(q.GetFloat(parameter.Limit, 10)),
		q.GetString(constant.Collection, ""),
		q.GetBool(constant.Full, false),
		q.GetString(constant.Mode, "hybrid"),
		metadata,
	)

	if outcome.Cause != nil {
		s.reporter.CaptureException(outcome.Cause)
	}

	return response.Success(notation.MarshalIndent(outcome.Results))
}
