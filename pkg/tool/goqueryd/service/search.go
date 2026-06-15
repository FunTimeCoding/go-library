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
	var outcome *store.SearchOutcome

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

		outcome = store.NewSearchOutcome(
			s.store.EnrichResults(results, metadata),
		)
	} else {
		o := store.NewSearchOption(query, limit)
		o.Collection = collection
		o.Full = full
		o.Metadata = metadata
		o.Reranker = s.reranker
		outcome = s.store.SearchWithFallback(o, s.ollama)
	}

	outcome.Facets = store.ComputeFacets(outcome.Results, 20)

	return outcome
}
