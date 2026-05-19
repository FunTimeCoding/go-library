package model_context

import (
	"context"
	"fmt"
	"github.com/funtimecoding/go-library/pkg/generative/mark/response"
	"github.com/funtimecoding/go-library/pkg/tool/goclauded/constant"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) edit(
	x context.Context,
	q mcp.CallToolRequest,
) (*mcp.CallToolResult, error) {
	identifier, e := q.RequireFloat(constant.Identifier)

	if e != nil {
		return response.Fail("identifier is required: %v", e)
	}

	message, f := q.RequireString(constant.Message)

	if f != nil {
		return response.Fail("message is required: %v", f)
	}

	if identifier < 0 {
		return response.Fail("invalid event identifier")
	}

	s.resolveCaller(x, constant.Edit)
	edited := s.service.Store.EditEvent(uint(identifier), message)

	switch edited.Kind {
	case constant.Summarize:
		s.service.Store.UpdateSummaryBody(edited.SessionIdentifier, message)
		s.pushSummary(edited.Name, message)
	case constant.Complete:
		s.service.Store.UpdateCompletionSummary(
			edited.SessionIdentifier,
			edited.Target,
			message,
		)
	}

	return response.Success(
		fmt.Sprintf("Updated event #%d", int(identifier)),
	)
}
