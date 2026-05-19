package base

import "time"

func (s *Server) Advance(d time.Duration) {
	*s.now = s.now.Add(d)
}
