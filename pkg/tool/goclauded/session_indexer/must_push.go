package session_indexer

import "github.com/funtimecoding/go-library/pkg/errors"

func (i *Indexer) MustPush(
	name string,
	body string,
	metadata map[string]string,
) {
	errors.PanicOnError(i.Push(name, body, metadata))
}
