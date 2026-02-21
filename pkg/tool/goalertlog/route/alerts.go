package route

import (
	"encoding/json"
	"github.com/funtimecoding/go-library/pkg/argument"
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/tool/goalertlog/store"
	"github.com/funtimecoding/go-library/pkg/web/constant"
	"net/http"
)

func Alerts(s *store.Store) http.HandlerFunc {
	return func(
		w http.ResponseWriter,
		r *http.Request,
	) {
		name := r.URL.Query().Get(argument.Name)

		if name == "" {
			http.Error(
				w,
				"missing name parameter",
				http.StatusBadRequest,
			)

			return
		}

		w.Header().Set(constant.ContentType, constant.Object)
		errors.PanicOnError(
			json.NewEncoder(w).Encode(toResponse(s.ByName(name))),
		)
	}
}
