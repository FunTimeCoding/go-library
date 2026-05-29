package service

import (
	"github.com/prometheus/client_golang/api/prometheus/v1"
	"github.com/prometheus/common/model"
)

func (s *Service) QueryRange(
	instance string,
	query string,
	r v1.Range,
) (model.Value, error) {
	c, e := s.Client(instance)

	if e != nil {
		return nil, e
	}

	var result model.Value
	e = safe(
		func() {
			result = c.QueryRange(query, r)
		},
	)

	return result, e
}
