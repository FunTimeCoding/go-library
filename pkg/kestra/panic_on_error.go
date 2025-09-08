package kestra

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/errors"
	"net/http"
)

func panicOnError(
	e error,
	r *http.Response,
) {
	if e != nil && r != nil && r.StatusCode >= 400 {
		fmt.Printf("Status: %d\n", r.StatusCode)
	}

	errors.PanicOnError(e)
}
