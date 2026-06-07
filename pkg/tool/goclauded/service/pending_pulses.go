package service

import "github.com/funtimecoding/go-library/pkg/tool/goclauded/store/pulse"

func (s *Service) PendingPulses(sessionIdentifier string) ([]pulse.Pulse, error) {
	return s.store.PendingPulses(sessionIdentifier)
}
