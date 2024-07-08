package status

import "github.com/funtimecoding/go-library/pkg/integers"

func (s *Status) Integer(v ...int) *Status {
	for _, element := range v {
		s.bubbles = append(s.bubbles, integers.ToString(element))
	}

	return s
}
