package site

func (s *Site) readMemories() string {
	s.protocol.WaitVisible(`//table//tbody/tr[1]`)

	return s.protocol.Outer("table")
}
