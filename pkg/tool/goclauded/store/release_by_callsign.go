package store

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/tool/goclauded/constant"
	"github.com/funtimecoding/go-library/pkg/tool/goclauded/store/session"
)

func (s *Store) ReleaseByCallsign(c string) {
	errors.PanicOnError(
		s.database.Model(session.Stub()).
			Where("callsign = ?", c).
			Update(constant.Callsign, nil).Error,
	)
}
