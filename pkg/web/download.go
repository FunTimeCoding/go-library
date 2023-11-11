package web

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/system"
	"log"
	"net/http"
)

func Download(locator string, output string) {
	f := system.Create(output)
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

	system.Copy(r.Body, f)
}
