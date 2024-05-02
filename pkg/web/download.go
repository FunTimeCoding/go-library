package web

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/system"
)

func Download(
	locator string,
	output string,
) {
	r := BasicGet(locator)
	defer func() {
		errors.LogClose(r.Body)
	}()

	errors.PanicStatus(r)
	f := system.Create(output)
	defer func() {
		errors.LogClose(f)
	}()

	system.Copy(r.Body, f)
}
