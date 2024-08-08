package status

import "github.com/funtimecoding/go-library/pkg/console/format"

func New(s *format.Settings) *Status {
	return &Status{
		linesByTag: make(map[string][]string),
		settings:   s,
	}
}
