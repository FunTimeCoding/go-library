package service

import "github.com/funtimecoding/go-library/pkg/tool/goclauded/constant"

func (s *Service) Summarize(
	sessionIdentifier string,
	name string,
	body string,
) error {
	if e := s.store.UpsertSummary(sessionIdentifier, name, body); e != nil {
		return e
	}

	if e := s.store.UpsertEvent(
		sessionIdentifier,
		constant.Summarize,
		name,
		"",
		body,
	); e != nil {
		return e
	}

	s.notify()

	return s.pushSummary(
		name,
		body,
		s.summaryMetadata(sessionIdentifier, name),
	)
}
