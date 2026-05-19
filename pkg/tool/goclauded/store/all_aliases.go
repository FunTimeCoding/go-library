package store

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/tool/goclauded/store/alias"
)

func (s *Store) AllAliases() map[string]string {
	var aliases []alias.Alias
	errors.PanicOnError(s.database.Find(&aliases).Error)
	result := make(map[string]string, len(aliases))

	for _, a := range aliases {
		result[a.SessionIdentifier] = a.Name
	}

	return result
}
