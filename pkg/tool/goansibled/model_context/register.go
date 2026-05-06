package model_context

import (
	"github.com/funtimecoding/go-library/pkg/tool/goansibled/constant"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) register() {
	s.server.AddTool(
		mcp.NewTool(
			constant.Playbooks,
			mcp.WithDescription(
				"List configured ansible playbook paths.",
			),
		),
		s.playbooks,
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.Trigger,
			mcp.WithDescription(
				"Trigger an ansible playbook run. Omit playbook to run all configured playbooks. Returns immediately; use runs/run tools to check results.",
			),
			mcp.WithString(
				constant.Playbook,
				mcp.Description(
					"Specific playbook path to run. Omit to run all.",
				),
			),
		),
		s.triggerRun,
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.Runs,
			mcp.WithDescription(
				"List recent ansible playbook runs. Returns metadata without full output.",
			),
			mcp.WithNumber(
				constant.Limit,
				mcp.Description("Maximum number of runs to return. Default: 20."),
			),
			mcp.WithString(
				constant.Status,
				mcp.Description(
					"Filter by status: pending, running, success, error.",
				),
			),
		),
		s.runs,
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.Run,
			mcp.WithDescription(
				"Fetch a single ansible playbook run by ID with full output.",
			),
			mcp.WithNumber(
				constant.Identifier,
				mcp.Required(),
				mcp.Description("Run ID."),
			),
		),
		s.getRun,
	)
}
