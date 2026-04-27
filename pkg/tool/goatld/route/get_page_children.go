package route

import (
	"github.com/funtimecoding/go-library/pkg/tool/goatld/convert"
	"github.com/funtimecoding/go-library/pkg/web"
	"net/http"
)

func (r *Router) GetPageChildren(
	w http.ResponseWriter,
	_ *http.Request,
	identifier string,
) {
	web.EncodeNotation(
		w,
		convert.ConfluencePagesFromPages(
			r.confluence.ChildPagesByIdentifier(identifier),
		),
	)
}
