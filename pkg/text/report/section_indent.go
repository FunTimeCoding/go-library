package report

func (s *Section) indentDeeper() {
	s.indent += 1

	for _, element := range s.renderables {
		element.indentDeeper()
	}
}
