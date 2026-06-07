package store

import (
	"github.com/funtimecoding/go-library/pkg/tool/goclauded/constant"
	"github.com/funtimecoding/go-library/pkg/tool/goclauded/store/event"
)

func (s *Store) EditEvent(
	identifier uint,
	body string,
) (*event.Event, error) {
	var existing event.Event

	if e := s.database.Where(
		"identifier = ?",
		identifier,
	).First(&existing).Error; e != nil {
		return nil, e
	}

	if e := s.database.Model(&existing).Update(
		constant.Body,
		body,
	).Error; e != nil {
		return nil, e
	}

	existing.Body = body

	return &existing, nil
}
