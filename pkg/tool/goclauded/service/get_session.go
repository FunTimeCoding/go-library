package service

import "github.com/funtimecoding/go-library/pkg/tool/goclauded/store/session"

func (s *Service) GetSession(identifier string) (*session.Session, error) {
	return s.store.GetSession(identifier)
}
