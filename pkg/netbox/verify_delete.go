package netbox

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/errors"
	"net/http"
)

func verifyDelete(
	entity string,
	r *http.Response,
	e error,
) {
	if r != nil {
		fmt.Printf("Delete %s code: %d\n", entity, r.StatusCode)
	}

	errors.PanicOnError(e)
}
