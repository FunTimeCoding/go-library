package service

import "github.com/funtimecoding/go-library/pkg/prometheus/alertmanager/alert"

func (s *Service) Alerts(instance string) ([]*alert.Alert, error) {
	c, e := s.Client(instance)

	if e != nil {
		return nil, e
	}

	var result []*alert.Alert
	e = safe(
		func() {
			result, _ = c.Alerts(nil)
		},
	)

	return result, e
}
