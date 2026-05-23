package store

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/tool/goclauded/constant"
	"github.com/funtimecoding/go-library/pkg/tool/goclauded/store/label"
)

func (s *Store) LabelsBySession(sessionIdentifier string) []label.Label {
	var result []label.Label
	errors.PanicOnError(
		s.database.Where(
			"session_identifier = ?",
			sessionIdentifier,
		).Order(constant.Key).Find(&result).Error,
	)

	return result
}
