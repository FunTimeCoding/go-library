package format

func (s *Settings) Color() *Settings {
	s.UseColor = true

	return s
}
