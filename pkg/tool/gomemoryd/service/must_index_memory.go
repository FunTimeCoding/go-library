package service

import "github.com/funtimecoding/go-library/pkg/errors"

func (s *Service) MustIndexMemory(
	path string,
	content string,
) {
	errors.PanicOnError(s.IndexMemory(path, content))
}
