package mock_indexer

import "github.com/funtimecoding/go-library/pkg/errors"

func (i *Indexer) MustPush(
	name string,
	body string,
) {
	errors.PanicOnError(i.Push(name, body))
}
