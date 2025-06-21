package site

func (s *Site) printMemories() {
	s.protocol.PrintNode(MemoriesSelector, []string{"class"})
}
