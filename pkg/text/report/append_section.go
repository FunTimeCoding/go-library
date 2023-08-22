package report

func (s *Section) AppendSection(other *Section) {
	other.indentDeeper()
	s.appendRenderable(other)
}
