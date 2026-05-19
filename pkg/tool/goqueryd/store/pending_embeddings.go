package store

import "github.com/funtimecoding/go-library/pkg/errors"

func (s *Store) pendingEmbeddings() []pendingDocument {
	rows, e := s.database.Query(
		`
		SELECT d.hash, c.body, MIN(d.path) AS path
		FROM document d
		JOIN content c ON d.hash = c.hash
		LEFT JOIN embedding e ON e.hash = d.hash AND e.sequence = 0
		WHERE d.active = 1 AND e.hash IS NULL
		GROUP BY d.hash
	`,
	)
	errors.PanicOnError(e)
	defer errors.PanicClose(rows)
	var result []pendingDocument

	for rows.Next() {
		var p pendingDocument
		errors.PanicOnError(rows.Scan(&p.hash, &p.body, &p.path))
		result = append(result, p)
	}

	errors.PanicOnError(rows.Err())

	return result
}
