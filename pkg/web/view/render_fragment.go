package view

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/web/constant"
	"maragu.dev/gomponents"
	"net/http"
)

func (v *View) RenderFragment(
	w http.ResponseWriter,
	fragment gomponents.Node,
) {
	w.Header().Set(constant.ContentType, constant.MarkupUnicode)
	errors.PanicOnError(fragment.Render(w))
}
