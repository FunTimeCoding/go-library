package status

func (s *Status) String(v ...string) *Status {
	for _, element := range v {
		s.bubbles = append(s.bubbles, element)
	}

	return s
}
