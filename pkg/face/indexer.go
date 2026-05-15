package face

type Indexer interface {
	Push(
		path string,
		body string,
	) error
	MustPush(
		path string,
		body string,
	)
	Delete(path string) error
	MustDelete(path string)
}
