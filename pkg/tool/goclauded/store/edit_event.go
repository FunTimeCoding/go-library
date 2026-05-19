package store

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/tool/goclauded/constant"
	"github.com/funtimecoding/go-library/pkg/tool/goclauded/store/event"
)

func (s *Store) EditEvent(
	identifier uint,
	body string,
) *event.Event {
	var existing event.Event
	errors.PanicOnError(
		s.database.Where("identifier = ?", identifier).First(&existing).Error,
	)
	errors.PanicOnError(
		s.database.Model(&existing).
			Update(constant.Body, body).Error,
	)
	s.notify()
	existing.Body = body

	return &existing
}
