package service

import (
	"github.com/funtimecoding/go-library/pkg/tool/goclauded/constant"
	"time"
)

func (s *Service) sweepInactivity() {
	cutoff := s.clock().Add(-1 * time.Hour)
	sessions := s.store.SweepInactivity(cutoff)

	for _, e := range sessions {
		s.store.LogEvent(
			e.Identifier,
			constant.InactivityTimeout,
			e.CallsignValue(),
			e.Topic,
			"",
		)
	}

	if len(sessions) > 0 {
		s.notify()
	}
}
