package status

import "github.com/funtimecoding/go-library/pkg/integers64"

func (s *Status) Integer64(v ...int64) *Status {
	for _, element := range v {
		s.bubbles = append(s.bubbles, integers64.ToString(element))
	}

	return s
}
