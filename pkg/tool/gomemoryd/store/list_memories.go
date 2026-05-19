package store

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"strings"
)

func (s *Store) ListMemories(
	memoryType string,
	tag string,
	activeOnly bool,
) ([]MemorySummary, error) {
	var parts []string
	var arguments []any
	parts = append(
		parts,
		`SELECT DISTINCT m.identifier, m.name, m.description, m.type, m.updated_at
		FROM memory m`,
	)

	if tag != "" {
		parts = append(
			parts,
			`JOIN memory_tag t ON t.memory_identifier = m.identifier AND t.tag = ?`,
		)
		arguments = append(arguments, tag)
	}

	parts = append(parts, `WHERE 1=1`)

	if activeOnly {
		parts = append(parts, `AND m.is_active = 1`)
	}

	if memoryType != "" {
		parts = append(parts, `AND m.type = ?`)
		arguments = append(arguments, memoryType)
	}

	parts = append(parts, `ORDER BY m.updated_at DESC`)
	rows, e := s.database.Query(strings.Join(parts, " "), arguments...)

	if e != nil {
		return nil, e
	}

	defer errors.LogClose(rows)
	var result []MemorySummary

	for rows.Next() {
		var m MemorySummary
		e := rows.Scan(
			&m.Identifier,
			&m.Name,
			&m.Description,
			&m.Type,
			&m.UpdatedAt,
		)

		if e != nil {
			return nil, e
		}

		m.Tags = s.tagsForMemory(m.Identifier)
		result = append(result, m)
	}

	return result, nil
}
