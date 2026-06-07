package service

import "github.com/funtimecoding/go-library/pkg/tool/goclauded/service/argument"

func (s *Service) EditSession(
	identifier string,
	a *argument.EditSession,
) error {
	if e := s.store.EditSession(identifier, a); e != nil {
		return e
	}

	s.notify()

	return nil
}
