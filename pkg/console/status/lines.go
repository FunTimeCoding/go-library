package status

func (s *Status) Lines(v ...string) *Status {
	for _, element := range v {
		if element == "" {
			continue
		}

		s.Line("  %s", element)
	}

	return s
}
