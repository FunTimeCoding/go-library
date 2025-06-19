package site

func (s *Site) printProfileButton() {
	s.protocol.PrintNode(ProfileButtonSelector, usefulAttributes)
}
