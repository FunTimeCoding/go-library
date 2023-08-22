package report

func (s *Section) appendRenderable(other renderable) {
	if s.maximumLength == NoLimit ||
		s.Length()+other.Length() <= s.maximumLength {
		s.renderables = append(s.renderables, other)
	}
}
