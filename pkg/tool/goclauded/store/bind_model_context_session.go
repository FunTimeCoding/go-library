package store

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/tool/goclauded/store/session"
)

func (s *Store) BindModelContextSession(
	name string,
	modelContextSessionIdentifier string,
) {
	errors.PanicOnError(
		s.database.Model(session.Stub()).
			Where("callsign = ?", name).
			Update(
				"model_context_session",
				modelContextSessionIdentifier,
			).Error,
	)
}
