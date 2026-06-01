package store

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/tool/gosproutd/store/seed"
)

func (s *Store) UpsertSeed(
	name string,
	path string,
	contentHash string,
	content string,
) {
	var existing seed.Seed
	result := s.mapper.Where("path = ?", path).Limit(1).Find(&existing)
	errors.PanicOnError(result.Error)

	if result.RowsAffected == 0 {
		var maxPosition int
		s.mapper.Model(&seed.Seed{}).Select("COALESCE(MAX(position), 0)").Scan(&maxPosition)
		errors.PanicOnError(s.mapper.Create(&seed.Seed{
			Name:        name,
			Path:        path,
			ContentHash: contentHash,
			Content:     content,
			Position:    maxPosition + 1,
		}).Error)

		return
	}

	if existing.ContentHash != contentHash {
		errors.PanicOnError(
			s.mapper.Model(&existing).Updates(map[string]any{
				"name":         name,
				"content_hash": contentHash,
				"content":      content,
			}).Error,
		)
	}
}
