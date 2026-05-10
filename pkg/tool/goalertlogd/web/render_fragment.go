package web

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/web/constant"
	"maragu.dev/gomponents"
	"net/http"
)

func renderFragment(
	w http.ResponseWriter,
	n gomponents.Node,
) {
	w.Header().Set(constant.ContentType, constant.MarkupUnicode)
	errors.PanicOnError(n.Render(w))
}
