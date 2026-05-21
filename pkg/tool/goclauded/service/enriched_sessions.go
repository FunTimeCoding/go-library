package service

import (
	"github.com/funtimecoding/go-library/pkg/tool/goclauded/service/enriched_session"
	"time"
)

func (s *Service) EnrichedSessions(
	limit int,
	offset int,
) []*enriched_session.Session {
	sessions := s.store.AllSessions(limit, offset)
	cutoff := s.clock().Add(-1 * time.Hour)
	result := make([]*enriched_session.Session, 0, len(sessions))

	for _, i := range sessions {
		e := enriched_session.New()
		e.Identifier = i.Identifier
		e.Slug = i.Slug
		e.Timestamp = i.SessionTimestamp
		e.CWD = i.CWD
		e.Branch = i.Branch
		e.Lines = i.Lines
		e.Name = i.Name
		e.Alias = i.AliasValue()
		e.Description = i.Description
		e.TurnCount = i.TurnCount
		e.Active = i.LastSeen.After(cutoff)
		result = append(result, e)
	}

	return result
}
