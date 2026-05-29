package service

import "time"

func (s *Service) LabelValues(
	instance string,
	label string,
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
			result = c.LabelValues(label, matches, since)
		},
	)

	return result, e
}
