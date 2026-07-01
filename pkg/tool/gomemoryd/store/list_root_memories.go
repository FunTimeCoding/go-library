package store

func (s *Store) ListRootMemories(
	memoryType string,
	tag string,
	activeOnly bool,
) ([]MemorySummary, error) {
	roots := true

	return s.queryMemories(memoryType, tag, activeOnly, &roots)
}
