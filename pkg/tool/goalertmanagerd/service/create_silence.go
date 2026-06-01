package service

import "time"

func (s *Service) CreateSilence(
	instance string,
	alert string,
	comment string,
	d time.Duration,
) (string, error) {
	c, e := s.Client(instance)

	if e != nil {
		return "", e
	}

	var result string
	e = safe(
		func() {
			result = c.SetSilence(alert, comment, d)
		},
	)

	return result, e
}
