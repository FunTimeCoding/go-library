package service

import "github.com/funtimecoding/go-library/pkg/tool/goclauded/constant"

func (s *Service) Moment(
	sessionIdentifier string,
	name string,
	line string,
) {
	s.store.LogEvent(
		sessionIdentifier,
		constant.Moment,
		name,
		"",
		line,
	)
	s.ForwardImpression(line, name)
	s.notify()
}
