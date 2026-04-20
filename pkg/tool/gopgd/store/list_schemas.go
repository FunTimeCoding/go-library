package store

import "context"

func (s *Store) ListSchemas(
	x context.Context,
	instance string,
) ([]map[string]any, error) {
	return s.Query(x, instance, listSchemasSQL)
}
