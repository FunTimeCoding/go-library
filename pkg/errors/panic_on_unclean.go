package errors

import (
	"errors"
	"net/http"
)

func PanicOnUnclean(e error) {
	if e == nil {
		return
	}

	if !errors.Is(e, http.ErrServerClosed) {
		panic(e)
	}
}
