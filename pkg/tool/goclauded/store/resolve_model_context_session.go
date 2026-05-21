package store

import "github.com/funtimecoding/go-library/pkg/tool/goclauded/store/session"

func (s *Store) ResolveModelContextSession(modelContextSessionIdentifier string) (string, string) {
	var i session.Session
	result := s.database.Where(
		"model_context_session = ?",
		modelContextSessionIdentifier,
	).Limit(1).Find(&i)

	if result.Error != nil || result.RowsAffected == 0 {
		return "", ""
	}

	return i.CallsignValue(), i.Identifier
}
