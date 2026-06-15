package store

func (s *Store) EnrichResults(
	results []SearchResult,
	metadata map[string]string,
) []SearchResult {
	var keys []documentKey

	for _, r := range results {
		keys = append(keys, documentKey{r.Collection, r.Path})
	}

	identifierMap := s.documentIdentifiers(keys)
	var identifiers []int

	for _, k := range keys {
		identifiers = append(identifiers, identifierMap[k])
	}

	s.enrichMetadata(results, identifiers)
	var enriched []SearchResult

	for _, r := range results {
		if !matchesMetadata(r.Metadata, r.SourceType, metadata) {
			continue
		}

		enriched = append(enriched, r)
	}

	return enriched
}
