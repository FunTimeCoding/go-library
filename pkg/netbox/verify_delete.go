package netbox

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/web"
	"net/http"
)

func verifyDelete(
	entity string,
	r *http.Response,
	e error,
) {
	if r != nil {
		b := web.ReadString(r)

		if b == "" {
			b = "empty body"
		}

		fmt.Printf("Delete %s (%d): %s\n", entity, r.StatusCode, b)
	}

	errors.PanicOnError(e)
}
