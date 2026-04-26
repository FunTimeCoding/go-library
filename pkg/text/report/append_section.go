package report

func (s *Section) AppendSection(other *Section) {
	other.IndentDeeper()
	s.appendRenderable(other)
}
