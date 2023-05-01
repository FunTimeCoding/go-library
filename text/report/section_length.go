package report

func (s *Section) Length() int {
	if len(s.renderables) == 0 {
		return len(s.title)
	}

	return len(s.Render())
}
