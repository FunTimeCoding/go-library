package store

import "database/sql"

func (s *Store) GetSourceType(
	collection string,
	pathPrefix string,
) string {
	pathPrefix = normalizePathPrefix(pathPrefix)
	var sourceType string
	e := s.database.QueryRow(
		"SELECT source_type FROM source_type_tag WHERE collection = ? AND path_prefix = ?",
		collection,
		pathPrefix,
	).Scan(&sourceType)

	if e == sql.ErrNoRows {
		return ""
	}

	return sourceType
}
