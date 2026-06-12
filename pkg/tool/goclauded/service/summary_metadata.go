package service

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/time"
)

func (s *Service) summaryMetadata(
	sessionIdentifier string,
	name string,
) map[string]string {
	result := map[string]string{
		"session_name": name,
		"session_id":   sessionIdentifier,
	}
	session, e := s.store.GetSession(sessionIdentifier)
	errors.PanicOnError(e)

	if session != nil && !session.StartedAt.IsZero() {
		result["date"] = session.StartedAt.Format(time.DateYear)
	}

	return result
}
