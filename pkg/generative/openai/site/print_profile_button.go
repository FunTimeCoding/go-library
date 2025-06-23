package site

func (s *Site) printProfile() {
	s.protocol.PrintNode(ProfileSelector, usefulAttributes)
}
