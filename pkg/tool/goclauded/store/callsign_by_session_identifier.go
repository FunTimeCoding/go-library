package store

import "github.com/funtimecoding/go-library/pkg/tool/goclauded/store/session"

func (s *Store) CallsignBySessionIdentifier(d string) string {
	var i session.Session
	result := s.database.Where(
		"identifier = ?",
		d,
	).Limit(1).Find(&i)

	if result.Error != nil || result.RowsAffected == 0 {
		return ""
	}

	return i.CallsignValue()
}
