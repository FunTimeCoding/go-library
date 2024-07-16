package format

func (s *Settings) Raw() *Settings {
	s.ShowRaw = true

	return s
}
