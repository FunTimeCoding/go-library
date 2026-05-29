package model_context

import (
	"context"
	"fmt"
	"github.com/funtimecoding/go-library/pkg/generative/mark/response"
	"github.com/funtimecoding/go-library/pkg/tool/goprometheusd/model_context/argument"
	"github.com/mark3labs/mcp-go/mcp"
	"strings"
	"time"
)

func (s *Server) labelValues(
	x context.Context,
	_ mcp.CallToolRequest,
	a argument.LabelValues,
) (*mcp.CallToolResult, error) {
	instance, okay := s.activeInstance(x)

	if !okay {
		return response.Fail(
			"no instance selected - use use_instance first",
		)
	}

	if a.Label == "" {
		return response.Fail("label is required")
	}

	var matches []string

	if a.Matches != "" {
		matches = strings.Split(a.Matches, ",")
	}

	since := time.Now().Add(-1 * time.Hour)

	if a.Since != "" {
		d, e := time.ParseDuration(a.Since)

		if e != nil {
			return response.Fail("invalid since: %s", e)
		}

		since = time.Now().Add(-d)
	}

	v, e := s.service.LabelValues(instance, a.Label, matches, since)

	if e != nil {
		return s.captureFail(
			e,
			fmt.Sprintf("label_values failed on %s", instance),
		)
	}

	return response.SuccessAny(v)
}
