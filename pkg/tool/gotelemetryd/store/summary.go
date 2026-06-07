package store

import "time"

func (s *Store) Summary(
	since string,
	until string,
	groupBy string,
) ([]SummaryRow, error) {
	selectClause := "tool, COUNT(*) as count"
	groupClause := "tool"

	if groupBy == "surface" {
		selectClause = "tool, surface, COUNT(*) as count"
		groupClause = "tool, surface"
	}

	if groupBy == "kind" {
		selectClause = "tool, kind, COUNT(*) as count"
		groupClause = "tool, kind"
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

	return result, query.Find(&result).Error
}
