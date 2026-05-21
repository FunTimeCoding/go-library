package service

import "github.com/funtimecoding/go-library/pkg/tool/goclauded/sweep"

func (s *Service) DeleteSession(identifier string) {
	s.client.Delete(identifier)
	sweep.DeleteSource(identifier)
	s.store.DeleteSession(identifier)
	s.notify()
}
