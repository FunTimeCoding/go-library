package web

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"maragu.dev/gomponents"
	"net/http"
)

func renderPage(
	w http.ResponseWriter,
	page gomponents.Node,
) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	errors.PanicOnError(page.Render(w))
}
