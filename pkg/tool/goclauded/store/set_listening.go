package store

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/tool/goclauded/store/session"
)

func (s *Store) SetListening(
	name string,
	listening bool,
) {
	errors.PanicOnError(
		s.database.Model(session.New()).
			Where("name = ?", name).
			Update("listening", listening).Error,
	)
	s.notify()
}
