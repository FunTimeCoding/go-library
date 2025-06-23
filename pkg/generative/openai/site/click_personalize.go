package site

func (s *Site) clickPersonalize() {
	s.protocol.ClickQuery(PersonalizeSelector)
}
