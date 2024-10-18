package output

func (s *Settings) Markdown() *Settings {
	s.Format = Markdown

	return s
}
