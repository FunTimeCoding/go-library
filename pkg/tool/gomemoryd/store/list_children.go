package store

func (s *Store) ListChildren(parentID int64) ([]MemorySummary, error) {
	return s.listMemoriesWithParent(parentID)
}
