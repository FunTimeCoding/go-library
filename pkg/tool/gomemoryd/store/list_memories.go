package store

func (s *Store) ListMemories(
	memoryType string,
	tag string,
	activeOnly bool,
) ([]MemorySummary, error) {
	return s.queryMemories(memoryType, tag, activeOnly, nil)
}
