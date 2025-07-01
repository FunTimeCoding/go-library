package status

func (s *Status) String(v ...string) *Status {
	for _, a := range v {
		if a == "" {
			continue
		}

		s.bubbles = append(s.bubbles, a)
	}

	return s
}
