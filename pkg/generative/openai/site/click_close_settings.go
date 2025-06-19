package site

func (s *Site) clickCloseSettings() {
	s.protocol.ClickQuery(CloseSettingsSelector)
}
