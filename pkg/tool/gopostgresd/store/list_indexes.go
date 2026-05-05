package store

import (
	"context"
	"fmt"
)

func (s *Store) ListIndexes(
	x context.Context,
	instance string,
	schema string,
	table string,
) ([]map[string]any, error) {
	return s.Query(
		x,
		instance,
		fmt.Sprintf(listIndexesSQL, schema, table),
	)
}
