package store

import "github.com/funtimecoding/go-library/pkg/generative/ollama"

func (s *Store) SearchWithFallback(
	o *SearchOption,
	l *ollama.Client,
) *SearchOutcome {
	status, e := s.Status()

	if e != nil {
		return &SearchOutcome{Degraded: true, Cause: e}
	}

	if status.TotalEmbeddings == 0 {
		results, f := s.SearchKeyword(o.Query, o.Limit, o.Collection, o.Full)

		if f != nil {
			return &SearchOutcome{Degraded: true, Cause: f}
		}

		return &SearchOutcome{Results: results, Degraded: true}
	}

	results, f := s.SearchHybrid(o, l)

	if f != nil {
		keyword, g := s.SearchKeyword(o.Query, o.Limit, o.Collection, o.Full)

		if g != nil {
			return &SearchOutcome{Degraded: true, Cause: g}
		}

		return &SearchOutcome{Results: keyword, Degraded: true, Cause: f}
	}

	return &SearchOutcome{Results: results}
}
