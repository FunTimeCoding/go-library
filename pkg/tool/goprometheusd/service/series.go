package service

import "github.com/prometheus/client_golang/api/prometheus/v1"

func (s *Service) Series(instance string) (v1.TSDBResult, error) {
	c, e := s.Client(instance)

	if e != nil {
		return v1.TSDBResult{}, e
	}

	var result v1.TSDBResult
	e = safe(
		func() {
			result = c.Series()
		},
	)

	return result, e
}
