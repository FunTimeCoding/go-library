package store

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/tool/goclauded/store/message"
)

func (s *Store) AllMessages(
	limit int,
	skip int,
) []message.Message {
	var result []message.Message
	errors.PanicOnError(
		s.database.
			Order("created_at DESC").
			Limit(limit).
			Offset(skip).
			Find(&result).Error,
	)

	return result
}
