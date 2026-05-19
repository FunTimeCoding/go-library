package store

func (s *Store) RunTimeoutSweep() {
	s.sweepInactivity()
	s.sweepCompleteTimeout()
}
