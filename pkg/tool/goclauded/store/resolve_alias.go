package store

import "github.com/funtimecoding/go-library/pkg/tool/goclauded/store/alias"

func (s *Store) ResolveAlias(name string) string {
	var a alias.Alias
	result := s.database.Where("name = ?", name).Limit(1).Find(&a)

	if result.Error != nil || result.RowsAffected == 0 {
		return ""
	}

	return a.SessionIdentifier
}
