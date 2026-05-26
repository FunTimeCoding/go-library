package service

import "github.com/funtimecoding/go-library/pkg/tool/goclauded/store/pulse"

func (s *Service) LatestPulse(sessionIdentifier string) *pulse.Pulse {
	return s.store.LatestPulse(sessionIdentifier)
}
