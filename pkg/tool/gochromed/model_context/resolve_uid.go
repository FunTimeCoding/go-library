package model_context

func (s *Server) resolveUID(
	tabID string,
	uid string,
) (int64, bool) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	cache, okay := s.snapshotCache[tabID]

	if !okay {
		return 0, false
	}

	v, okay := cache[uid]

	return v, okay
}
