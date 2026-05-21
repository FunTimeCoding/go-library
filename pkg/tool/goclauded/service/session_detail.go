package service

import "github.com/funtimecoding/go-library/pkg/tool/goclauded/service/session_detail"

func (s *Service) SessionDetail(query string) *session_detail.Detail {
	r := s.store.ResolveSessionIdentifier(query)

	if !r.Found() {
		return nil
	}

	e := s.store.GetSession(r.Identifier())

	if e == nil {
		return nil
	}

	result := session_detail.New()
	result.Identifier = e.Identifier
	result.Slug = e.Slug
	result.Created = e.SessionTimestamp
	result.Session = e
	result.Alias = e.AliasValue()
	result.Description = e.Description
	result.TurnCount = e.TurnCount
	result.Completions = s.store.CompletionsBySession(e.Identifier)
	result.Summary = s.store.SummaryBySession(e.Identifier)

	return result
}
