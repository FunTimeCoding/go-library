package route

import (
	"encoding/json"
	"github.com/funtimecoding/go-library/pkg/argument"
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/tool/goalertlog/store"
	"github.com/funtimecoding/go-library/pkg/web/constant"
	"net/http"
	"time"
)

func Recent(s *store.Store) http.HandlerFunc {
	return func(
		w http.ResponseWriter,
		r *http.Request,
	) {
		now := time.Now()
		start := now.Add(-1 * time.Hour)
		end := now
		startParam := r.URL.Query().Get(argument.Start)

		if startParam != "" {
			parsed, e := time.Parse(DateFormat, startParam)

			if e != nil {
				http.Error(
					w,
					"invalid start parameter",
					http.StatusBadRequest,
				)

				return
			}

			start = parsed
		}

		endParam := r.URL.Query().Get(argument.End)

		if endParam != "" {
			parsed, e := time.Parse(DateFormat, endParam)

			if e != nil {
				http.Error(
					w,
					"invalid end parameter",
					http.StatusBadRequest,
				)

				return
			}

			end = parsed
		}

		w.Header().Set(constant.ContentType, constant.Object)
		errors.PanicOnError(
			json.NewEncoder(w).Encode(toResponse(s.ByTimeRange(start, end))),
		)
	}
}
