package gorilla

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/web/location"
	"net/http"
)

func home(
	w http.ResponseWriter,
	r *http.Request,
) {
	errors.PanicOnError(
		homeTemplate.Execute(
			w,
			fmt.Sprintf("ws://%s%s", r.Host, location.Echo),
		),
	)
}
