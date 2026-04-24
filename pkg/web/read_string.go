package web

import "net/http"

func ReadString(r *http.Response) string {
	return string(ReadBytes(r))
}
