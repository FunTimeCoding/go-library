package service

import (
	"github.com/prometheus/common/model"
	"time"
)

func (s *Service) Query(
	instance string,
	query string,
	t time.Time,
) (model.Value, error) {
	c, e := s.Client(instance)

	if e != nil {
		return nil, e
	}

	var result model.Value
	e = safe(
		func() {
			result = c.Query(query, t)
		},
	)

	return result, e
}
