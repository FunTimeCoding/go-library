package mock_indexer

func (i *Indexer) Push(
	name string,
	body string,
) error {
	i.Pushed = append(i.Pushed, PushCall{Name: name, Body: body})

	return nil
}
