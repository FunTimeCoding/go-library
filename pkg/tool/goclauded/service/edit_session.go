package service

import "github.com/funtimecoding/go-library/pkg/tool/goclauded/service/argument"

func (s *Service) EditSession(
	identifier string,
	a *argument.EditSession,
) {
	s.store.EditSession(identifier, a)
	s.notify()
}
