package store

func (s *Store) ResolveSourceType(
	collection string,
	documentPath string,
) string {
	prefixes := sourceTypePrefixes(documentPath)

	for _, prefix := range prefixes {
		var sourceType string
		e := s.database.QueryRow(
			`SELECT source_type FROM source_type_tag
			WHERE (collection = ? OR collection = '') AND path_prefix = ?
			ORDER BY CASE WHEN collection = ? THEN 0 ELSE 1 END
			LIMIT 1`,
			collection,
			prefix,
			collection,
		).Scan(&sourceType)

		if e == nil {
			return sourceType
		}
	}

	return ""
}
