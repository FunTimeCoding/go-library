package store

import (
	"context"
	"fmt"
)

func (s *Store) ListTables(
	x context.Context,
	instance string,
	schema string,
) ([]map[string]any, error) {
	return s.Query(
		x,
		instance,
		fmt.Sprintf(listTablesSQL, schema),
	)
}
