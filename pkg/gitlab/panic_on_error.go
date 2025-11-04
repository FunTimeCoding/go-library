package gitlab

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"gitlab.com/gitlab-org/api/client-go"
)

func panicOnError(
	r *gitlab.Response,
	e error,
) {
	if r == nil {
		errors.PanicOnError(e)

		return
	}

	errors.PanicOnWebError(r.Response, e)
}
