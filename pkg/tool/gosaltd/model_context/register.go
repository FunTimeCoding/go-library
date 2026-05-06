package model_context

import (
	"github.com/funtimecoding/go-library/pkg/tool/gosaltd/constant"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) register() {
	s.server.AddTool(
		mcp.NewTool(
			constant.Trigger,
			mcp.WithDescription(
				"Trigger a salt highstate on all minions. Returns immediately; use runs/run tools to check results.",
			),
		),
		s.triggerRun,
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.Runs,
			mcp.WithDescription(
				"List recent salt highstate runs. Returns metadata without full output.",
			),
			mcp.WithNumber(
				constant.Limit,
				mcp.Description(
					"Maximum number of runs to return. Default: 20.",
				),
			),
		),
		s.runs,
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.Run,
			mcp.WithDescription(
				"Fetch a single salt highstate run by ID with full output.",
			),
			mcp.WithNumber(
				constant.Identifier,
				mcp.Required(),
				mcp.Description("Run ID."),
			),
		),
		s.getRun,
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.Minions,
			mcp.WithDescription(
				"List connected salt minions.",
			),
		),
		s.minions,
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.Keys,
			mcp.WithDescription(
				"List salt minion key status: accepted, pending, denied, rejected.",
			),
		),
		s.keys,
	)
}
