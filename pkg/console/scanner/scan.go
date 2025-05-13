package scanner

import "github.com/funtimecoding/go-library/pkg/errors"

func (s *Scanner) Scan() []string {
	var result []string

	for s.scanner.Scan() {
		result = append(result, s.scanner.Text())
	}

	errors.PanicOnError(s.scanner.Err())

	return result
}
