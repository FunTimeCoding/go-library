package report

func (s *Section) IndentDeeper() {
	s.indent += 1

	for _, e := range s.renderables {
		e.IndentDeeper()
	}
}
