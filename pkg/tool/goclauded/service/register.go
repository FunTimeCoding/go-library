package service

import "github.com/funtimecoding/go-library/pkg/tool/goclauded/store/ensure_result"

func (s *Service) Register(sessionIdentifier string) *ensure_result.Result {
	result := s.store.EnsureSession(sessionIdentifier)
	s.store.MarkNeedsRoster(sessionIdentifier)
	s.EnrichSession(sessionIdentifier)
	s.store.LogEvent(
		sessionIdentifier,
		"register",
		result.Callsign,
		"",
		"",
	)
	s.notify()

	return result
}
