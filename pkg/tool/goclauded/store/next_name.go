package store

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/tool/goclauded/constant"
	"github.com/funtimecoding/go-library/pkg/tool/goclauded/store/session"
	"math/rand"
)

func (s *Store) NextName() string {
	var taken []string
	errors.PanicOnError(
		s.database.Model(session.New()).Pluck(
			constant.SessionName,
			&taken,
		).Error,
	)
	takenSet := map[string]bool{}

	for _, name := range taken {
		takenSet[name] = true
	}

	var available []string

	for _, name := range s.poolNames() {
		if !takenSet[name] {
			available = append(available, name)
		}
	}

	if len(available) == 0 {
		return ""
	}

	return available[rand.Intn(len(available))]
}
