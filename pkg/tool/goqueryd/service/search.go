package service

import "github.com/funtimecoding/go-library/pkg/tool/goqueryd/store"

func (s *Service) Search(
	query string,
	limit int,
	collection string,
	full bool,
	sourceType string,
	mode string,
) *store.SearchOutcome {
	if mode == "keyword" {
		results, e := s.Store.SearchKeyword(
			query,
			limit,
			collection,
			full,
		)

		if e != nil {
			return store.NewDegradedSearchOutcome(e)
		}

		return store.NewSearchOutcome(results)
	}

	o := store.NewSearchOption(query, limit)
	o.Collection = collection
	o.Full = full
	o.SourceType = sourceType
	o.Reranker = s.reranker

	return s.Store.SearchWithFallback(o, s.ollama)
}
