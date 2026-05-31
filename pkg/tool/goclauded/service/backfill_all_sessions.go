package service

func (s *Service) BackfillAllSessions() *BackfillResult {
	r := &BackfillResult{}

	for _, e := range s.store.AllSessions(0, 0) {
		resolved := s.client.Resolve(e.Identifier)

		if resolved.Identifier == "" {
			r.Skipped++

			continue
		}

		s.EnrichSession(e.Identifier)
		r.Enriched++
	}

	return r
}
