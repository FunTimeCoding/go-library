package route

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"net/http"
)

func (h *Router) DeleteEntry(
	w http.ResponseWriter,
	_ *http.Request,
	id int,
) {
	errors.PanicOnError(h.store.Delete(uint(id)))
	w.WriteHeader(http.StatusNoContent)
}
