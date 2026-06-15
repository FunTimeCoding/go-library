package service

import "github.com/funtimecoding/go-library/pkg/tool/gomemoryd/store"

func (s *Service) SearchRelevant(
	query string,
	limit int,
) ([]store.SearchResult, error) {
	results, e := s.searcher.Search(query, "memories", limit)

	if e != nil {
		return nil, e
	}

	var matches []store.SearchResult

	for _, r := range results {
		identifier, e := extractIdentifier(r.Path)

		if e != nil {
			continue
		}

		m, e := s.store.GetMemory(identifier)

		if e != nil {
			continue
		}

		matches = append(
			matches,
			store.SearchResult{
				Identifier:  m.Identifier,
				Name:        m.Name,
				Content:     m.Content,
				Description: m.Description,
				Type:        m.Type,
				UpdatedAt:   m.UpdatedAt,
				Rank:        r.Score,
				Tags:        m.Tags,
			},
		)
	}

	return matches, nil
}
