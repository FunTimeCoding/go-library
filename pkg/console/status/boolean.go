package status

import "github.com/funtimecoding/go-library/pkg/booleans"

func (s *Status) Boolean(v ...bool) *Status {
	for _, element := range v {
		s.bubbles = append(s.bubbles, booleans.ToString(element))
	}

	return s
}
