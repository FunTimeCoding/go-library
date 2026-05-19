package store

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/tool/goclauded/constant"
	"github.com/funtimecoding/go-library/pkg/tool/goclauded/store/summary"
)

func (s *Store) UpdateSummaryBody(
	sessionIdentifier string,
	body string,
) {
	errors.PanicOnError(
		s.database.Model(summary.Stub()).
			Where("session_identifier = ?", sessionIdentifier).
			Update(constant.Body, body).Error,
	)
	s.notify()
}
