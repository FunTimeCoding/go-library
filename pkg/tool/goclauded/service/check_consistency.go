package service

import "github.com/funtimecoding/go-library/pkg/tool/goclauded/constant"

func (s *Service) CheckConsistency() {
	databaseSessions := s.store.AllSessions(0, 0)
	databaseSet := make(map[string]bool, len(databaseSessions))

	for _, e := range databaseSessions {
		databaseSet[e.Identifier] = true
	}

	fileSessions := s.client.Sessions()
	fileSet := make(map[string]bool, len(fileSessions))

	for _, e := range fileSessions {
		fileSet[e.Identifier] = true
	}

	for _, e := range fileSessions {
		if !databaseSet[e.Identifier] {
			s.store.CreateDiscovered(e.Identifier)
			s.EnrichSession(e.Identifier)
			s.logger.Structured(
				"consistency_discovered",
				constant.Identifier,
				e.Identifier,
				constant.Slug,
				e.Slug,
			)
		}
	}

	for _, e := range databaseSessions {
		if !fileSet[e.Identifier] {
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
