package route

import (
	"encoding/json"
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/tool/goatld/convert"
	"github.com/funtimecoding/go-library/pkg/tool/goatld/server"
	"github.com/funtimecoding/go-library/pkg/web"
	"net/http"
)

func (h *Router) CreatePage(
	w http.ResponseWriter,
	q *http.Request,
) {
	var b server.CreatePageJSONRequestBody
	errors.PanicOnError(json.NewDecoder(q.Body).Decode(&b))
	web.ObjectHeader(w)
	w.WriteHeader(http.StatusCreated)
	web.Encode(
		w,
		convert.ConfluencePageDetail(
			h.confluence.CreatePage(
				b.SpaceIdentifier,
				b.ParentIdentifier,
				b.Title,
				b.Body,
			),
		),
	)
}
