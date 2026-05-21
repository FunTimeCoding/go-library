package service

import "github.com/funtimecoding/go-library/pkg/generative/anthropic/claude/message"

func (s *Service) Messages(sessionIdentifier string) []message.Message {
	return s.client.Messages(sessionIdentifier)
}
