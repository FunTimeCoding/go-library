package store

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/tool/goclauded/store/message"
)

func (s *Store) SendMessage(
	fromName string,
	toName string,
	body string,
) {
	errors.PanicOnError(
		s.database.Create(message.New(fromName, toName, body)).Error,
	)
	s.notify()
}
