package store

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/tool/goqueryd/constant"
)

func (s *Store) AddCollection(
	name string,
	path string,
	pattern string,
) {
	if pattern == "" {
		pattern = constant.DefaultGlob
	}

	_, e := s.database.Exec(
		`INSERT INTO collection (name, path, pattern) VALUES (?, ?, ?)
		ON CONFLICT(name) DO UPDATE SET path = excluded.path, pattern = excluded.pattern`,
		name,
		path,
		pattern,
	)
	errors.PanicOnError(e)
}
