package model_context

import (
	"context"
	"github.com/funtimecoding/go-library/pkg/chromium/snapshot"
	"github.com/funtimecoding/go-library/pkg/generative/mark/response"
	"github.com/funtimecoding/go-library/pkg/tool/gochromed/constant"
	"github.com/funtimecoding/go-library/pkg/tool/gochromed/model_context/argument"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) Fill(
	_ context.Context,
	_ mcp.CallToolRequest,
	a argument.Fill,
) (*mcp.CallToolResult, error) {
	if a.UID == "" {
		return response.Fail("uid is required")
	}

	if a.Value == "" {
		return response.Fail("value is required")
	}

	t, e := s.resolveTab(a.TabID, a.Title, a.URL)

	if e != nil {
		return response.Fail(e.Error())
	}

	backendID, okay := s.resolveUID(t.Identifier, a.UID)

	if !okay {
		return response.Fail(
			"uid %s not found - take a snapshot first",
			a.UID,
		)
	}

	x := s.client.AcquireTarget(t.Identifier)
	e = withTimeoutAction(
		constant.TargetTimeout,
		func() error {
			return s.client.FillNode(x, backendID, a.Value)
		},
	)

	if e != nil {
		return s.captureDetail(e)
	}

	if a.Snapshot != nil && !*a.Snapshot {
		return response.Success("filled %s", a.UID)
	}

	nodes, e := withTimeout(
		constant.TargetTimeout,
		func() ([]*snapshot.Node, error) {
			return s.client.Snapshot(x)
		},
	)

	if e != nil {
		return s.captureDetail(e)
	}

	s.cacheSnapshot(t.Identifier, nodes)

	return response.Success(
		"filled %s\n\n%s",
		a.UID,
		snapshot.Format(nodes, 0),
	)
}
