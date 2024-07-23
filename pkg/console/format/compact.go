package format

func (s *Settings) Compact() *Settings {
	s.UseCompact = true

	return s
}
