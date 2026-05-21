package service

import "github.com/funtimecoding/go-library/pkg/generative/anthropic/claude/session"

func (s *Service) Sessions() []*session.Session {
	return s.client.Sessions()
}
