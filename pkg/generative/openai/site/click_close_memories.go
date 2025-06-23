package site

func (s *Site) clickCloseMemories() {
	n := s.protocol.Select(CloseMemoriesSelector, 2)

	if n == nil {
		return
	}

	s.protocol.ClickSearch(n.FullXPath())
}
