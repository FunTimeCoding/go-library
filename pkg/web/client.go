package web

import "net/http"

func Client() *http.Client {
	return &http.Client{}
}
