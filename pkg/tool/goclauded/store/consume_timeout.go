package store

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/tool/goclauded/constant"
	"github.com/funtimecoding/go-library/pkg/tool/goclauded/store/session"
)

func (s *Store) ConsumeTimeout(name string) string {
	var i session.Session
	result := s.database.Where("callsign = ?", name).Limit(1).Find(&i)

	if result.Error != nil || result.RowsAffected == 0 || i.TimedOut == "" {
		return ""
	}

	reason := i.TimedOut
	errors.PanicOnError(
		s.database.Model(session.Stub()).
			Where("callsign = ?", name).
			Update("timed_out", "").Error,
	)

	switch reason {
	case constant.InactivityTimeout:
		return fmt.Sprintf(
			"You were timed out due to 1 hour inactivity. Last topic: %s",
			i.Topic,
		)
	case constant.CompleteTimeout:
		return "You were timed out after completing without re-announcing."
	default:
		return ""
	}
}
