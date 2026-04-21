package route

import (
	"github.com/funtimecoding/go-library/pkg/tool/gonetbd/server"
	"github.com/funtimecoding/go-library/pkg/web"
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

	web.EncodeNotation(w, result)
}
