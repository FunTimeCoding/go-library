package route

import (
	"github.com/funtimecoding/go-library/pkg/tool/gonetbd/server"
	"github.com/funtimecoding/go-library/pkg/web"
	"net/http"
)

func (h *Router) ListClusterTypes(
	w http.ResponseWriter,
	_ *http.Request,
) {
	types := h.client.ClusterTypes()
	result := make([]server.ClusterType, 0, len(types))

	for _, t := range types {
		result = append(
			result,
			server.ClusterType{Identifier: t.Identifier, Name: t.Name},
		)
	}

	web.EncodeNotation(w, result)
}
