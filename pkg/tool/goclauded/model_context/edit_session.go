package model_context

import (
	"context"
	"github.com/funtimecoding/go-library/pkg/generative/mark/response"
	"github.com/funtimecoding/go-library/pkg/tool/goclauded/constant"
	"github.com/funtimecoding/go-library/pkg/tool/goclauded/service/argument"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) editSession(
	x context.Context,
	q mcp.CallToolRequest,
) (*mcp.CallToolResult, error) {
	c := s.resolveCaller(x, constant.EditSession)
	target := q.GetString(constant.Target, "")
	var sessionIdentifier string

	if target != "" {
		r := s.service.ResolveSessionIdentifier(target)

		if r.Ambiguous() {
			return response.Fail(
				"ambiguous: %q matches %d sessions",
				target,
				len(r.Matches),
			)
		}

		if !r.Found() {
			return response.Fail("session not found: %s", target)
		}

		sessionIdentifier = r.Identifier()
	} else if c.Callsign != "" {
		sessionIdentifier = c.SessionIdentifier
	}

	if sessionIdentifier == "" {
		return response.Fail(
			"could not resolve session - announce first or specify a target",
		)
	}

	a := argument.NewEditSession()
	hasChange := false

	if v := q.GetString(constant.Alias, ""); v != "" {
		a.Alias = &v
		hasChange = true
	}

	if v := q.GetString(constant.Description, ""); v != "" {
		a.Description = &v
		hasChange = true
	}

	if v := q.GetString(constant.Topic, ""); v != "" {
		a.Topic = &v
		hasChange = true
	}

	if v := q.GetString(constant.Files, ""); v != "" {
		a.Files = &v
		hasChange = true
	}

	if !hasChange {
		return response.Fail("at least one field required")
	}

	s.service.EditSession(sessionIdentifier, a)

	return response.Success("Session updated")
}
