package model_context

import (
	"context"
	library "github.com/funtimecoding/go-library/pkg/constant"
	"github.com/funtimecoding/go-library/pkg/generative/mark/response"
	"github.com/funtimecoding/go-library/pkg/strings/join"
	"github.com/funtimecoding/go-library/pkg/tool/goclauded/constant"
	"github.com/funtimecoding/go-library/pkg/tool/goclauded/service/argument"
	"github.com/funtimecoding/go-library/pkg/tool/goclauded/timeline"
	"github.com/mark3labs/mcp-go/mcp"
	"time"
)

func (s *Server) history(
	x context.Context,
	q mcp.CallToolRequest,
) (*mcp.CallToolResult, error) {
	s.resolveCaller(x, constant.History)
	a := argument.Timeline{
		Limit:  int(q.GetFloat(constant.Limit, 20)),
		Offset: int(q.GetFloat(constant.Offset, 0)),
		Kind:   q.GetString(constant.Kind, ""),
		Full:   q.GetBool(constant.Full, false),
	}
	sinceRaw := q.GetString(constant.Since, "")
	beforeRaw := q.GetString(constant.Before, "")

	if sinceRaw != "" {
		sinceDuration, e := time.ParseDuration(sinceRaw)

		if e != nil {
			return response.Fail("invalid since duration: %v", e)
		}

		a.Since = time.Now().Add(-sinceDuration).UTC().Format(time.RFC3339)
	}

	if beforeRaw != "" {
		beforeDuration, e := time.ParseDuration(beforeRaw)

		if e != nil {
			return response.Fail("invalid before duration: %v", e)
		}

		a.Before = time.Now().Add(-beforeDuration).UTC().Format(time.RFC3339)
	}

	merged, f := s.service.Timeline(a)

	if f != nil {
		return s.captureFail(f, library.UnexpectedError)
	}

	if len(merged) == 0 {
		return response.Success("No events recorded.")
	}

	var lines []string

	for _, e := range merged {
		lines = append(lines, timeline.Format(e))
	}

	return response.Success(join.NewLine(lines))
}
