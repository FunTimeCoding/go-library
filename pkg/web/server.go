package web

import (
	"net/http"
	"time"
)

func Server(
	h http.Handler,
	address string,
) *http.Server {
	return &http.Server{
		Addr:         address,
		Handler:      h,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
}
