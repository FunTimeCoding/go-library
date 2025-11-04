package confluence

import (
	"github.com/ctreminiom/go-atlassian/v2/pkg/infra/models"
	"github.com/funtimecoding/go-library/pkg/errors"
)

func panicOnError(
	r *models.ResponseScheme,
	e error,
) {
	errors.PanicOnWebError(r.Response, e)
}
