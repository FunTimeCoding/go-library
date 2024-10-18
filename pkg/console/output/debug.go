package output

func (s *Settings) Debug() *Settings {
	s.ShowDebug = true

	return s
}
