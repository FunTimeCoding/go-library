package service

import "time"

func (s *Service) LabelNames(
	instance string,
	matches []string,
	since time.Time,
) ([]string, error) {
	c, e := s.Client(instance)

	if e != nil {
		return nil, e
	}

	var result []string
	e = safe(
		func() {
			result = c.LabelNames(matches, since)
		},
	)

	return result, e
}
