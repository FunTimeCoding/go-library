package store

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/tool/goclauded/store/completion"
)

func (s *Store) CreateCompletion(
	sessionIdentifier string,
	name string,
	kind string,
	topic string,
	summary string,
) {
	errors.PanicOnError(
		s.database.Create(
			completion.New(
				sessionIdentifier,
				name,
				kind,
				topic,
				summary,
			),
		).Error,
	)
}
