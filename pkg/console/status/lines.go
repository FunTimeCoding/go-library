package status

func (s *Status) Lines(v ...string) *Status {
	for _, e := range v {
		if e == "" {
			continue
		}

		s.Line("  %s", e)
	}

	return s
}
