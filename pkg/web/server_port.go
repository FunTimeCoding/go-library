package web

import (
	"fmt"
	"net/http"
)

func ServerPort(
	h http.Handler,
	port int,
) *http.Server {
	return Server(h, fmt.Sprintf(":%d", port))
}
