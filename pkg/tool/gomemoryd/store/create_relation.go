package store

import "time"

func (s *Store) CreateRelation(
	sourceID int64,
	targetID int64,
) error {
	now := time.Now().UTC().Format(time.RFC3339)
	_, e := s.database.Exec(
		`INSERT OR IGNORE INTO memory_relation (source_identifier, target_identifier, created_at)
		VALUES (?, ?, ?)`,
		sourceID,
		targetID,
		now,
	)

	return e
}
