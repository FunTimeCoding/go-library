package service

func (s *Service) ReconcileSummaries() {
	for _, m := range s.store.ListSummaries() {
		s.indexer.MustPush(m.Name, m.Body)
	}
}
