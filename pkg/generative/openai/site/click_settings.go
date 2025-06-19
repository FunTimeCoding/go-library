package site

func (s *Site) clickSettings() {
	s.protocol.ClickQuery(SettingsButtonSelector)
}
