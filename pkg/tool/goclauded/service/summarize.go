package service

import "github.com/funtimecoding/go-library/pkg/tool/goclauded/constant"

func (s *Service) Summarize(
	sessionIdentifier string,
	name string,
	body string,
) error {
	s.store.UpsertSummary(sessionIdentifier, name, body)
	s.store.UpsertEvent(
		sessionIdentifier,
		constant.Summarize,
		name,
		"",
		body,
	)
	s.notify()

	return s.pushSummary(
		name,
		body,
		s.summaryMetadata(sessionIdentifier, name),
	)
}
