package route

import (
	"encoding/json"
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/tool/gopgd/server"
	"github.com/funtimecoding/go-library/pkg/web/constant"
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

	w.Header().Set(constant.ContentType, constant.Object)
	errors.PanicOnError(json.NewEncoder(w).Encode(result))
}
