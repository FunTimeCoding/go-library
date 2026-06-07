package virtual_file_system

import "github.com/funtimecoding/go-library/pkg/errors"

func (s *System) MustReadDirectory(path string) []string {
	result, e := s.ReadDirectory(path)
	errors.PanicOnError(e)

	return result
}
