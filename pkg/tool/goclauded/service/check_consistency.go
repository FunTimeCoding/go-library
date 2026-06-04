package service

import "github.com/funtimecoding/go-library/pkg/tool/goclauded/constant"

func (s *Service) CheckConsistency() {
	databaseSessions := s.store.AllSessions(0, 0)
	databaseSet := make(map[string]bool, len(databaseSessions))

	for _, e := range databaseSessions {
		databaseSet[e.Identifier] = true
	}

	s.statesMu.Lock()
	cacheSet := make(map[string]bool, len(s.states))

	for identifier := range s.states {
		cacheSet[identifier] = true
	}

	s.statesMu.Unlock()

	for identifier := range cacheSet {
		if !databaseSet[identifier] {
			s.store.CreateDiscovered(identifier)
			s.RefreshFromCache(identifier)
			s.logger.Structured(
				"consistency_discovered",
				constant.Identifier,
				identifier,
			)
		}
	}

	for _, e := range databaseSessions {
		if !cacheSet[e.Identifier] {
			s.logger.Structured(
				"consistency_missing_jsonl",
				constant.Identifier,
				e.Identifier,
				constant.SessionName,
				e.Name,
				constant.Alias,
				e.Alias,
			)
		}
	}
}
