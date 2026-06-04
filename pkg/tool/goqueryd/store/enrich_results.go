package store

import "github.com/funtimecoding/go-library/pkg/tool/goqueryd/constant"

func (s *Store) EnrichResults(
	results []SearchResult,
	metadata map[string]string,
) []SearchResult {
	var enriched []SearchResult

	for _, r := range results {
		identifier := s.documentIdentifier(r.Collection, r.Path)
		r.Metadata = s.GetMetadata(identifier)

		if r.Metadata != nil {
			r.SourceType = r.Metadata[constant.SourceType]
		}

		if r.SourceType == "" {
			r.SourceType = s.ResolveSourceType(r.Collection, r.Path)
		}

		if !matchesMetadata(r.Metadata, r.SourceType, metadata) {
			continue
		}

		enriched = append(enriched, r)
	}

	return enriched
}
