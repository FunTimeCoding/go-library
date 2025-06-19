package site

func (s *Site) clickCloseMemories() {
	n := s.protocol.Select(CloseDialogSelector, 2)

	if n == nil {
		return
	}

	s.protocol.ClickSearch(n.FullXPath())
}
