package web

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	system2 "github.com/funtimecoding/go-library/pkg/system"
	"log"
	"net/http"
)

func Download(locator string, output string) {
	f := system2.Create(output)
	defer func() {
		errors.LogClose(f)
	}()

	r := BasicGet(locator)
	defer func() {
		errors.LogClose(r.Body)
	}()

	if r.StatusCode != http.StatusOK {
		log.Panicf(
			"unexpected status code: %d",
			r.StatusCode,
		)
	}

	system2.Copy(r.Body, f)
}
