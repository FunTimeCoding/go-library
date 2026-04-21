package route

import (
	"github.com/funtimecoding/go-library/pkg/tool/gonetbd/server"
	"github.com/funtimecoding/go-library/pkg/web"
	"net/http"
)

func (h *Router) ListClusters(
	w http.ResponseWriter,
	_ *http.Request,
) {
	clusters := h.client.Clusters()
	result := make([]server.Cluster, 0, len(clusters))

	for _, cl := range clusters {
		entry := server.Cluster{Identifier: cl.Identifier, Name: cl.Name}

		if cl.Raw.Type.Name != "" {
			entry.Type = &cl.Raw.Type.Name
		}

		result = append(result, entry)
	}

	web.EncodeNotation(w, result)
}
