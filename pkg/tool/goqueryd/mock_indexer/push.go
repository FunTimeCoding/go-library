package mock_indexer

func (i *Indexer) Push(
	name string,
	body string,
	metadata map[string]string,
) error {
	i.Pushed = append(
		i.Pushed,
		PushCall{Name: name, Body: body, Metadata: metadata},
	)

	return nil
}
