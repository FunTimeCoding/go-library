package errors

import (
	"fmt"
	"net/http"
)

func PanicOnWebError(
	r *http.Response,
	e error,
) {
	if e != nil && r != nil && r.StatusCode >= 400 {
		fmt.Printf("Status: %d\n", r.StatusCode)
		fmt.Printf("Error: %s\n", e.Error())

		// Using web.PrintHeader would create a circular dependency
		for k, v := range r.Header {
			fmt.Printf("Header: %s: %s\n", k, v)
		}

		fmt.Printf("Response: %+v\n", r)
	}

	PanicOnError(e)
}
