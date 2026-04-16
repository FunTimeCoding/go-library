package route

import (
	"encoding/json"
	"github.com/funtimecoding/go-library/pkg/errors"
	generated "github.com/funtimecoding/go-library/pkg/tool/gonetbd/server"
	"github.com/funtimecoding/go-library/pkg/web/constant"
	"net/http"
)

func (h *Router) CreateSite(
	w http.ResponseWriter,
	q *http.Request,
) {
	var body generated.CreateNameRequest
	errors.PanicOnError(json.NewDecoder(q.Body).Decode(&body))
	s := h.client.CreateSite(body.Name)
	w.Header().Set(constant.ContentType, constant.Object)
	w.WriteHeader(http.StatusCreated)
	errors.PanicOnError(
		json.NewEncoder(w).Encode(
			generated.Site{
				Identifier: s.Identifier, Name: s.Name,
			},
		),
	)
}
