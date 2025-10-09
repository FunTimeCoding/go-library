package web

import (
	"fmt"
	"net/http"
)

func PrintHeader(h http.Header) {
	for k, v := range h {
		fmt.Printf("Header: %s: %s\n", k, v)
	}
}
