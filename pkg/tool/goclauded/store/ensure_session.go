package store

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/tool/goclauded/store/ensure_result"
	"github.com/funtimecoding/go-library/pkg/tool/goclauded/store/session"
)

func (s *Store) EnsureSession(identifier string) *ensure_result.Result {
	var i session.Session
	result := s.database.Where("identifier = ?", identifier).Limit(1).Find(&i)
	errors.PanicOnError(result.Error)

	if result.RowsAffected > 0 {
		previousLastSeen := i.LastSeen
		s.database.Model(&i).Updates(
			map[string]any{
				"last_seen":  s.clock(),
				"turn_count": i.TurnCount + 1,
			},
		)

		return ensure_result.NewExisting(i.CallsignValue(), previousLastSeen)
	}

	name := s.NextName()

	if name == "" {
		return ensure_result.Stub()
	}

	errors.PanicOnError(
		s.database.Create(session.NewRegistered(identifier, name)).Error,
	)

	return ensure_result.New(name)
}
