package mock_indexer

type PushCall struct {
	Name     string
	Body     string
	Metadata map[string]string
}
