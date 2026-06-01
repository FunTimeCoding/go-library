package service

import "github.com/funtimecoding/go-library/pkg/prometheus/alertmanager/silence"

func (s *Service) Silences(
	instance string,
	expired bool,
) ([]*silence.Silence, error) {
	c, e := s.Client(instance)

	if e != nil {
		return nil, e
	}

	var result []*silence.Silence
	e = safe(
		func() {
			result = c.Silences(expired)
		},
	)

	return result, e
}
