package status

func (s *Status) String(v ...string) *Status {
	s.bubbles = append(s.bubbles, v...)

	return s
}
