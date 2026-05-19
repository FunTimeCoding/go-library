package mock_indexer

func (i *Indexer) Delete(path string) error {
	i.Deleted = append(i.Deleted, path)

	return nil
}
