package store

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/tool/goclauded/store/message"
)

func (s *Store) PendingMessages(name string) []message.Message {
	var result []message.Message
	errors.PanicOnError(
		s.database.
			Where("(to_name = ? OR to_name = '') AND read = ?", name, false).
			Order("created_at").
			Find(&result).Error,
	)

	if len(result) > 0 {
		errors.PanicOnError(
			s.database.Model(message.Stub()).
				Where(
					"(to_name = ? OR to_name = '') AND read = ?",
					name,
					false,
				).
				Update("read", true).Error,
		)
	}

	return result
}
