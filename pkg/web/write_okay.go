package web

import "net/http"

func WriteOkay(
	w http.ResponseWriter,
	s string,
) int {
	return Write(w, http.StatusOK, s)
}
