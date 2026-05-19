package store

import "github.com/funtimecoding/go-library/pkg/tool/goclauded/store/session"

func (s *Store) ResolveSessionIdentifier(query string) string {
	var i session.Session

	if r := s.database.Where(
		"name = ?",
		query,
	).Limit(1).Find(&i); r.RowsAffected > 0 {
		return i.Identifier
	}

	if r := s.database.Where(
		"identifier = ?",
		query,
	).Limit(1).Find(&i); r.RowsAffected > 0 {
		return i.Identifier
	}

	if identifier := s.ResolveAlias(query); identifier != "" {
		return identifier
	}

	return ""
}
