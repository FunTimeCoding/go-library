package service

func (s *Service) DeleteSilence(
	instance string,
	identifier string,
) error {
	c, e := s.Client(instance)

	if e != nil {
		return e
	}

	return safe(
		func() {
			c.DeleteSilence(identifier)
		},
	)
}
