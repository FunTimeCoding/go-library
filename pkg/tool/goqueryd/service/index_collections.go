package service

import "github.com/funtimecoding/go-library/pkg/tool/goqueryd/service/index_result"

func (s *Service) IndexCollections(collection string) ([]*index_result.IndexResult, error) {
	status, e := s.Store.Status()

	if e != nil {
		return nil, e
	}

	var results []*index_result.IndexResult

	for _, c := range status.Collections {
		if collection != "" && c.Name != collection {
			continue
		}

		r := s.Store.Index(c.Name)
		results = append(
			results,
			index_result.New(
				c.Name,
				r.Indexed,
				r.Updated,
				r.Unchanged,
				r.Removed,
			),
		)
	}

	return results, nil
}
