package store

import (
	"github.com/funtimecoding/go-library/pkg/tool/goclauded/constant"
	"github.com/funtimecoding/go-library/pkg/tool/goclauded/store/pool_name"
)

func (s *Store) poolNames() []string {
	var entries []pool_name.PoolName
	s.database.Order(constant.Identifier).Find(&entries)

	if len(entries) == 0 {
		return defaultNames
	}

	result := make([]string, len(entries))

	for i, e := range entries {
		result[i] = e.Name
	}

	return result
}
