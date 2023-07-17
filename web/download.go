package web

import (
	"fmt"
	"github.com/funtimecoding/go-library/errors"
	"io"
	"net/http"
	"os"
)

func Download(locator string, output string) error {
	file, fileFail := os.Create(output)
	errors.PanicOnError(fileFail)
	defer func() {
		errors.LogOnError(file.Close())
	}()

	response, getFail := http.Get(locator)
	errors.PanicOnError(getFail)
	defer func() {
		errors.LogOnError(response.Body.Close())
	}()

	if response.StatusCode != http.StatusOK {
		return fmt.Errorf(
			"unexpected status code: %d",
			response.StatusCode,
		)
	}

	_, copyFail := io.Copy(file, response.Body)
	errors.PanicOnError(copyFail)

	return nil
}
