package claude

import "github.com/funtimecoding/go-library/pkg/generative/anthropic/claude/message"

type ToolContextResult struct {
	ToolName string
	ToolID   string
	Before   []message.Message
	After    []message.Message
}
