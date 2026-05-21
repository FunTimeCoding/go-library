package service

import "github.com/funtimecoding/go-library/pkg/tool/goclauded/store/session"

func (s *Service) GetSession(identifier string) *session.Session {
	return s.store.GetSession(identifier)
}
