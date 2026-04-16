package route

import (
	"encoding/json"
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/tool/gonetbd/server"
	"github.com/funtimecoding/go-library/pkg/web/constant"
	"net/http"
)

func (h *Router) ListTags(
	w http.ResponseWriter,
	_ *http.Request,
) {
	tags := h.client.Tags()
	result := make([]server.Tag, 0, len(tags))

	for _, t := range tags {
		result = append(
			result,
			server.Tag{Identifier: t.Identifier, Name: t.Name},
		)
	}

	w.Header().Set(constant.ContentType, constant.Object)
	errors.PanicOnError(json.NewEncoder(w).Encode(result))
}
