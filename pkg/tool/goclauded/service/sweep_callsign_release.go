package service

import "time"

func (s *Service) sweepCallsignRelease() {
	cutoff := s.clock().Add(-72 * time.Hour)
	sessions := s.store.SweepCallsignRelease(cutoff)

	for _, e := range sessions {
		s.logger.Structured(
			"callsign_released",
			"callsign",
			e.CallsignValue(),
			"identifier",
			e.Identifier,
		)
	}
}
