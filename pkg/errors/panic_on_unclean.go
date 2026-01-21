package errors

import "net/http"

func PanicOnUnclean(e error) {
	if e == nil {
		return
	}

	if !Is(e, http.ErrServerClosed) {
		panic(e)
	}
}
