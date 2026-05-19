package store

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/tool/goclauded/constant"
	"github.com/funtimecoding/go-library/pkg/tool/goclauded/store/alias"
)

func (s *Store) SetDescription(
	sessionIdentifier string,
	description string,
) {
	var a alias.Alias
	result := s.database.Where(
		"session_identifier = ?",
		sessionIdentifier,
	).Limit(1).Find(&a)
	errors.PanicOnError(result.Error)

	if result.RowsAffected == 0 {
		errors.PanicOnError(
			s.database.Create(
				alias.NewDescription(sessionIdentifier, description),
			).Error,
		)
	} else {
		errors.PanicOnError(
			s.database.Model(alias.Stub()).
				Where("session_identifier = ?", sessionIdentifier).
				Update(constant.Description, description).Error,
		)
	}

	s.notify()
}
