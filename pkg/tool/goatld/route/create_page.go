package route

import (
	"encoding/json"
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/tool/goatld/convert"
	"github.com/funtimecoding/go-library/pkg/tool/goatld/server"
	"github.com/funtimecoding/go-library/pkg/web"
	"net/http"
)

func (r *Router) CreatePage(
	w http.ResponseWriter,
	e *http.Request,
) {
	var s server.CreatePageJSONRequestBody
	errors.PanicOnError(json.NewDecoder(e.Body).Decode(&s))
	web.ObjectHeader(w)
	w.WriteHeader(http.StatusCreated)
	web.Encode(
		w,
		convert.ConfluencePageDetail(
			r.confluence.CreatePage(
				s.SpaceIdentifier,
				s.ParentIdentifier,
				s.Title,
				s.Body,
			),
		),
	)
}
