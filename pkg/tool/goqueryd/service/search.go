package service

import "github.com/funtimecoding/go-library/pkg/tool/goqueryd/store"

func (s *Service) Search(
	query string,
	limit int,
	collection string,
	full bool,
	mode string,
	metadata map[string]string,
) *store.SearchOutcome {
	if mode == "keyword" {
		results, e := s.store.SearchKeyword(
			query,
			limit,
			collection,
			full,
		)

		if e != nil {
			return store.NewDegradedSearchOutcome(e)
		}

		return store.NewSearchOutcome(
			s.store.EnrichResults(results, metadata),
		)
	}

	o := store.NewSearchOption(query, limit)
	o.Collection = collection
	o.Full = full
	o.Metadata = metadata
	o.Reranker = s.reranker

	return s.store.SearchWithFallback(o, s.ollama)
}
