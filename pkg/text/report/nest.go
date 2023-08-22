package report

func (s *Section) Nest(
	title string,
	maximumLength int,
) *Section {
	nested := &Section{
		title:         title,
		indent:        s.indent + 1,
		maximumLength: maximumLength,
	}
	s.appendRenderable(nested)

	return nested
}
