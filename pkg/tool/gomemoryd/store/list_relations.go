package store

import "github.com/funtimecoding/go-library/pkg/errors"

func (s *Store) ListRelations(identifier int64) ([]Relation, error) {
	rows, e := s.database.Query(
		`SELECT source_identifier, target_identifier, created_at FROM memory_relation
		WHERE source_identifier = ? OR target_identifier = ?
		ORDER BY created_at`,
		identifier,
		identifier,
	)

	if e != nil {
		return nil, e
	}

	defer errors.LogClose(rows)
	var result []Relation

	for rows.Next() {
		var r Relation
		e := rows.Scan(&r.SourceIdentifier, &r.TargetIdentifier, &r.CreatedAt)

		if e != nil {
			return nil, e
		}

		result = append(result, r)
	}

	return result, nil
}
