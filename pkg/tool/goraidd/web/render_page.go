package web

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/web/constant"
	g "maragu.dev/gomponents"
	"net/http"
)

func renderPage(
	w http.ResponseWriter,
	page g.Node,
) {
	w.Header().Set(constant.ContentType, "text/html; charset=utf-8")
	errors.PanicOnError(page.Render(w))
}
