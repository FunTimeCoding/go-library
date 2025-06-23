package site

func (s *Site) clickMemories() {
	n := s.protocol.Select(MemoriesSelector, 1)

	if n == nil {
		return
	}

	s.protocol.ClickSearch(n.FullXPath())
}
