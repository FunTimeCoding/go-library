package summary_indexer

import "github.com/funtimecoding/go-library/pkg/errors"

func (i *Indexer) MustDelete(path string) {
	errors.PanicOnError(i.Delete(path))
}
