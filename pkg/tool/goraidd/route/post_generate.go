package route

import "net/http"

func (h *Router) PostGenerate(
	w http.ResponseWriter,
	_ *http.Request,
) {
	w.WriteHeader(http.StatusNotImplemented)
}
