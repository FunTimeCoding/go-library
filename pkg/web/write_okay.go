package web

import "net/http"

func WriteOkay(w http.ResponseWriter, s string) {
	Write(w, http.StatusOK, s)
}
