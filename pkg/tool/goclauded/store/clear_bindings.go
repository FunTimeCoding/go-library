package store

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/tool/goclauded/store/session"
)

func (s *Store) ClearBindings() {
	errors.PanicOnError(
		s.database.Model(session.New()).
			Where("model_context_session != ''").
			Updates(
				map[string]any{
					"model_context_session": "",
					"needs_reannounce":      true,
				},
			).Error,
	)
}
