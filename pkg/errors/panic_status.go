package errors

import (
	"log"
	"net/http"
)

func PanicStatus(r *http.Response) {
	if r.StatusCode != http.StatusOK {
		log.Panicf(
			"unexpected status code: %d",
			r.StatusCode,
		)
	}
}
