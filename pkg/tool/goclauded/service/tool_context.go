package service

import "github.com/funtimecoding/go-library/pkg/generative/anthropic/claude"

func (s *Service) ToolContext(
	sessionIdentifier string,
	toolFilter string,
	surroundCount int,
) []claude.ToolContextResult {
	return s.client.ToolContext(sessionIdentifier, toolFilter, surroundCount)
}
