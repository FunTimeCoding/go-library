package store

import "time"

func (s *Store) CreateImpression(
	content string,
	source string,
) (int64, error) {
	now := time.Now().UTC().Format(time.RFC3339)
	result, e := s.database.Exec(
		`INSERT INTO impression (content, source, created_at) VALUES (?, ?, ?)`,
		content,
		source,
		now,
	)

	if e != nil {
		return 0, e
	}

	return result.LastInsertId()
}
