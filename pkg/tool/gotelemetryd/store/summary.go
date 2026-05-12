package store

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"time"
)

func (s *Store) Summary(
	since string,
	until string,
	groupBy string,
) []SummaryRow {
	selectClause := "tool, COUNT(*) as count"
	groupClause := "tool"

	if groupBy == "surface" {
		selectClause = "tool, surface, COUNT(*) as count"
		groupClause = "tool, surface"
	}

	query := s.mapper.Model(&UsageEvent{}).
		Select(selectClause).
		Group(groupClause).
		Order("count DESC")

	if since != "" {
		t, e := time.Parse(time.RFC3339, since)

		if e == nil {
			query = query.Where("created_at >= ?", t)
		}
	}

	if until != "" {
		t, e := time.Parse(time.RFC3339, until)

		if e == nil {
			query = query.Where("created_at <= ?", t)
		}
	}

	var result []SummaryRow
	errors.PanicOnError(query.Find(&result).Error)

	return result
}
