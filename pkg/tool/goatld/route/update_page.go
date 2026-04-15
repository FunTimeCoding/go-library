package route

import (
	"encoding/json"
	"github.com/funtimecoding/go-library/pkg/errors"
	generated "github.com/funtimecoding/go-library/pkg/tool/goatld/server"
	"github.com/funtimecoding/go-library/pkg/web/constant"
	"net/http"
)

func (h *Router) UpdatePage(
	w http.ResponseWriter,
	q *http.Request,
	identifier string,
) {
	var body generated.UpdatePageJSONRequestBody
	errors.PanicOnError(json.NewDecoder(q.Body).Decode(&body))
	message := ""

	if body.Message != nil {
		message = *body.Message
	}

	page := h.confluence.UpdatePage(
		identifier,
		body.Title,
		body.Body,
		message,
	)
	w.Header().Set(constant.ContentType, constant.Object)
	errors.PanicOnError(
		json.NewEncoder(w).Encode(toConfluencePageDetail(page)),
	)
}
