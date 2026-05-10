package web

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/web/constant"
	"maragu.dev/gomponents"
	"net/http"
)

func renderPage(
	w http.ResponseWriter,
	page gomponents.Node,
) {
	w.Header().Set(constant.ContentType, constant.MarkupUnicode)
	errors.PanicOnError(page.Render(w))
}
