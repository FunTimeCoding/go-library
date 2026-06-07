package store

import "github.com/funtimecoding/go-library/pkg/tool/goclauded/store/session"

func (s *Store) DeleteSession(identifier string) error {
	return s.database.Where(
		"identifier = ?",
		identifier,
	).Delete(session.Stub()).Error
}
