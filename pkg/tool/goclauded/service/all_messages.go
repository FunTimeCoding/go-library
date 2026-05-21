package service

import "github.com/funtimecoding/go-library/pkg/tool/goclauded/store/message"

func (s *Service) AllMessages(
	limit int,
	skip int,
) []message.Message {
	return s.store.AllMessages(limit, skip)
}
