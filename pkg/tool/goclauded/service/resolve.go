package service

import "github.com/funtimecoding/go-library/pkg/generative/anthropic/claude/session"

func (s *Service) Resolve(query string) *session.Session {
	return s.client.Resolve(query)
}
