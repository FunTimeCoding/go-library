package model_context

import (
	"context"
	"fmt"
	"github.com/funtimecoding/go-library/pkg/generative/mark/response"
	"github.com/funtimecoding/go-library/pkg/tool/goclauded/constant"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) label(
	x context.Context,
	q mcp.CallToolRequest,
) (*mcp.CallToolResult, error) {
	c := s.resolveCaller(x, constant.Label)
	key, e := q.RequireString(constant.Key)

	if e != nil {
		return response.Fail("key is required: %v", e)
	}

	target := q.GetString(constant.Target, "")
	var sessionIdentifier string
	var targetName string

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
		targetName = target
	} else if c.Callsign != "" {
		sessionIdentifier = c.SessionIdentifier
	}

	if sessionIdentifier == "" {
		return response.Fail(
			"could not resolve session - announce first or specify a target",
		)
	}

	arguments := q.GetArguments()
	raw, hasValue := arguments[constant.Value]

	if !hasValue {
		value := s.service.GetLabel(sessionIdentifier, key)

		if value == "" {
			return response.Success("no label %s", key)
		}

		return response.Success("%s=%s", key, value)
	}

	value, _ := raw.(string)

	if value == "" {
		change := s.service.DeleteLabel(
			sessionIdentifier,
			c.Callsign,
			targetName,
			key,
		)

		return response.Success(change)
	}

	change := s.service.SetLabel(
		sessionIdentifier,
		c.Callsign,
		targetName,
		key,
		value,
	)

	return response.Success(
		fmt.Sprintf("%s on %s", change, displayName(c, target)),
	)
}
