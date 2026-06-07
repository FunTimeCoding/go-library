package service

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/tool/goclauded/constant"
	"time"
)

func (s *Service) sweepCompleteTimeout() {
	cutoff := s.clock().Add(-30 * time.Minute)
	sessions := s.store.SweepCompleteTimeout(cutoff)

	for _, e := range sessions {
		errors.PanicOnError(
			s.store.LogEvent(
				e.Identifier,
				constant.CompleteTimeout,
				e.CallsignValue(),
				"",
				"",
			),
		)
	}

	if len(sessions) > 0 {
		s.notify()
	}
}
