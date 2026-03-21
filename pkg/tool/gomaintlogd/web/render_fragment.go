package web

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/web/constant"
	g "maragu.dev/gomponents"
	"net/http"
)

func renderFragment(
	w http.ResponseWriter,
	fragment g.Node,
) {
	w.Header().Set(constant.ContentType, "text/html; charset=utf-8")
	errors.PanicOnError(fragment.Render(w))
}
