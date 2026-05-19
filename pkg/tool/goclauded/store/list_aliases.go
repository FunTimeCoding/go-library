package store

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/tool/goclauded/constant"
	"github.com/funtimecoding/go-library/pkg/tool/goclauded/store/alias"
)

func (s *Store) ListAliases() []alias.Alias {
	var result []alias.Alias
	errors.PanicOnError(
		s.database.Order(constant.SessionName).Find(&result).Error,
	)

	return result
}
