package store

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/tool/goclauded/constant"
	"github.com/funtimecoding/go-library/pkg/tool/goclauded/service/argument"
	"github.com/funtimecoding/go-library/pkg/tool/goclauded/store/session"
)

func (s *Store) EditSession(
	identifier string,
	a *argument.EditSession,
) {
	updates := map[string]any{}

	if a.Alias != nil {
		if *a.Alias == "" {
			updates[constant.Alias] = nil
		} else {
			updates[constant.Alias] = *a.Alias
		}
	}

	if a.Description != nil {
		updates[constant.Description] = *a.Description
	}

	if a.Topic != nil {
		updates[constant.Topic] = *a.Topic
	}

	if a.Files != nil {
		updates[constant.Files] = *a.Files
	}

	if len(updates) == 0 {
		return
	}

	errors.PanicOnError(
		s.database.Model(session.Stub()).
			Where("identifier = ?", identifier).
			Updates(updates).Error,
	)
}
