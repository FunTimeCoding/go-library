package service

import "github.com/funtimecoding/go-library/pkg/tool/goclauded/store/ensure_result"

func (s *Service) Register(sessionIdentifier string) (*ensure_result.Result, error) {
	result, e := s.store.EnsureSession(sessionIdentifier)

	if e != nil {
		return nil, e
	}

	if e := s.store.MarkNeedsRoster(sessionIdentifier); e != nil {
		return nil, e
	}

	s.EnrichSession(sessionIdentifier)

	if e := s.store.LogEvent(
		sessionIdentifier,
		"register",
		result.Callsign,
		"",
		"",
	); e != nil {
		return nil, e
	}

	s.notify()

	return result, nil
}
