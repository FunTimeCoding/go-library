package site

func (s *Site) readMemories() string {
	if !s.protocol.HasNodes(`//table`) {
		return ""
	}

	s.protocol.WaitVisible(`//table//tbody/div[1]`)

	return s.protocol.Outer("table")
}
