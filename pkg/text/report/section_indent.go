package report

func (s *Section) indentDeeper() {
	s.indent += 1

	for _, e := range s.renderables {
		e.indentDeeper()
	}
}
