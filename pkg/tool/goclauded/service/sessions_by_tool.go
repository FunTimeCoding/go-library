package service

import "github.com/funtimecoding/go-library/pkg/generative/anthropic/claude"

func (s *Service) SessionsByTool(toolFilter string) []*claude.SessionToolCount {
	return s.client.SessionsByTool(toolFilter)
}
