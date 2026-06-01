package model_context

import (
	"context"
	"fmt"
	"github.com/funtimecoding/go-library/pkg/generative/mark/response"
	"github.com/funtimecoding/go-library/pkg/tool/goalertmanagerd/model_context/argument"
	"github.com/mark3labs/mcp-go/mcp"
	"time"
)

func (s *Server) createSilence(
	x context.Context,
	_ mcp.CallToolRequest,
	a argument.CreateSilence,
) (*mcp.CallToolResult, error) {
	instance, okay := s.activeInstance(x)

	if !okay {
		return response.Fail(
			"no instance selected - use use_instance first",
		)
	}

	if a.Alert == "" {
		return response.Fail("alert is required")
	}

	var d time.Duration

	if a.Duration != "" {
		var e error
		d, e = time.ParseDuration(a.Duration)

		if e != nil {
			return response.Fail("invalid duration: %s", e)
		}
	}

	v, e := s.service.CreateSilence(instance, a.Alert, a.Comment, d)

	if e != nil {
		return s.captureFail(
			e,
			fmt.Sprintf("create_silence failed on %s", instance),
		)
	}

	return response.Success("silence created: %s", v)
}
