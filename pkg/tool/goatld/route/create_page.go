package route

import (
	"encoding/json"
	"github.com/funtimecoding/go-library/pkg/errors"
	generated "github.com/funtimecoding/go-library/pkg/tool/goatld/server"
	"github.com/funtimecoding/go-library/pkg/web/constant"
	"net/http"
)

func (h *Router) CreatePage(
	w http.ResponseWriter,
	q *http.Request,
) {
	var body generated.CreatePageJSONRequestBody
	errors.PanicOnError(json.NewDecoder(q.Body).Decode(&body))
	page := h.confluence.CreatePage(
		body.SpaceIdentifier,
		body.ParentIdentifier,
		body.Title,
		body.Body,
	)
	w.Header().Set(constant.ContentType, constant.Object)
	w.WriteHeader(http.StatusCreated)
	errors.PanicOnError(
		json.NewEncoder(w).Encode(toConfluencePageDetail(page)),
	)
}
