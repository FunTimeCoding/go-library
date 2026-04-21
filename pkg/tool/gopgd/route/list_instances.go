package route

import (
	"github.com/funtimecoding/go-library/pkg/tool/gopgd/server"
	"github.com/funtimecoding/go-library/pkg/web"
	"net/http"
)

func (h *Router) ListInstances(
	w http.ResponseWriter,
	_ *http.Request,
) {
	var result []server.Instance

	for _, i := range h.store.Instances() {
		result = append(
			result,
			server.Instance{
				Name:     i.Name,
				Host:     i.Host,
				Port:     i.Port,
				Database: i.Database,
			},
		)
	}

	web.EncodeNotation(w, result)
}
