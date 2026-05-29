package worker

import "github.com/funtimecoding/go-library/pkg/errors"

func (w *Worker) poll() {
	_, e := w.service.IndexCollections("")
	errors.PanicOnError(e)
	_, e = w.service.Embed()
	errors.PanicOnError(e)
}
