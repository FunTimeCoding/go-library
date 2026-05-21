package service

import "github.com/funtimecoding/go-library/pkg/errors"

func (s *Service) MustSummarize(
	sessionIdentifier string,
	name string,
	body string,
) {
	errors.PanicOnError(s.Summarize(sessionIdentifier, name, body))
}
