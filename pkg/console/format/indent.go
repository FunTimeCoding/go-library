package format

func (s *Settings) Indent(i int) *Settings {
	s.Indentation = i

	return s
}
