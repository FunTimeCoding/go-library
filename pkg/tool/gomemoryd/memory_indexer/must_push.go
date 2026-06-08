package memory_indexer

import "github.com/funtimecoding/go-library/pkg/errors"

func (i *Indexer) MustPush(
	path string,
	body string,
	metadata map[string]string,
) {
	errors.PanicOnError(i.Push(path, body, metadata))
}
