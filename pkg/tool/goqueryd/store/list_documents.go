package store

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/strings/join"
	"github.com/funtimecoding/go-library/pkg/tool/goqueryd/store/result"
)

func (s *Store) ListDocuments(collection string) ([]result.DocumentEntry, error) {
	rows, e := s.database.Query(
		`SELECT d.collection, d.path, d.title
		FROM document d
		WHERE d.collection = ? AND d.active = 1
		ORDER BY d.path`,
		collection,
	)

	if e != nil {
		return nil, e
	}

	defer errors.PanicClose(rows)
	var r []result.DocumentEntry

	for rows.Next() {
		var c, p, title string

		if f := rows.Scan(&c, &p, &title); f != nil {
			return nil, f
		}

		r = append(
			r,
			result.DocumentEntry{
				VirtualPath: buildVirtualPath(c, p),
				FilePath:    join.Empty(c, "/", p),
				Title:       title,
			},
		)
	}

	if e := rows.Err(); e != nil {
		return nil, e
	}

	return r, nil
}
