package site

func (s *Site) printCloseSettings() {
	s.protocol.PrintNode(CloseSettingsSelector, usefulAttributes)
}
