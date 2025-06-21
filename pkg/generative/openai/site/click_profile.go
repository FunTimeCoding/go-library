package site

func (s *Site) clickProfile() {
	s.protocol.ClickQuery(ProfileSelector)
}
