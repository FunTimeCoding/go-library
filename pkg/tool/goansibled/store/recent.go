package store

import "time"

func (s *Store) Recent(limit int) ([]PlaybookRun, error) {
	var result []PlaybookRun

	return result, s.mapper.
		Where("created_at > ?", time.Now().Add(-RetentionAge)).
		Order("created_at desc").
		Limit(limit).
		Find(&result).Error
}
