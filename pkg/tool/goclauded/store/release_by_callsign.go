package store

import (
	"github.com/funtimecoding/go-library/pkg/tool/goclauded/constant"
	"github.com/funtimecoding/go-library/pkg/tool/goclauded/store/session"
)

func (s *Store) ReleaseByCallsign(c string) error {
	return s.database.Model(session.Stub()).Where(
		"callsign = ?",
		c,
	).Update(constant.Callsign, nil).Error
}
