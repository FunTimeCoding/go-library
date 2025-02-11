package gorilla

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/errors"
	"net/http"
)

func home(
	w http.ResponseWriter,
	r *http.Request,
) {
	errors.PanicOnError(
		homeTemplate.Execute(w, fmt.Sprintf("ws://%s/echo", r.Host)),
	)
}
