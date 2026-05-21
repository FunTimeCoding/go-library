package service

import "github.com/funtimecoding/go-library/pkg/tool/goclauded/store/session"

func (s *Service) AllSessions(
	limit int,
	offset int,
) []session.Session {
	return s.store.AllSessions(limit, offset)
}
