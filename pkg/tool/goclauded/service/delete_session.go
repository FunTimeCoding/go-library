package service

import "github.com/funtimecoding/go-library/pkg/tool/goclauded/sweep"

func (s *Service) DeleteSession(identifier string) error {
	s.client.Delete(identifier)
	sweep.DeleteSource(identifier)

	if e := s.store.DeleteSession(identifier); e != nil {
		return e
	}

	s.notify()

	return nil
}
