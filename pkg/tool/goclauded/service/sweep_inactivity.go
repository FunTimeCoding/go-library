package service

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/tool/goclauded/constant"
	"time"
)

func (s *Service) sweepInactivity() {
	cutoff := s.clock().Add(-1 * time.Hour)
	sessions := s.store.SweepInactivity(cutoff)

	for _, e := range sessions {
		errors.PanicOnError(
			s.store.LogEvent(
				e.Identifier,
				constant.InactivityTimeout,
				e.CallsignValue(),
				map[string]string{"topic": e.Topic},
			),
		)
	}

	if len(sessions) > 0 {
		s.notify()
	}
}
