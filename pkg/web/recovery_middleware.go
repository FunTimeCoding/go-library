package web

import (
	"github.com/funtimecoding/go-library/pkg/web/constant"
	"github.com/getsentry/sentry-go"
	"net/http"
)

func RecoveryMiddleware(h *sentry.Hub) func(http.Handler) http.Handler {
	return func(n http.Handler) http.Handler {
		return http.HandlerFunc(
			func(
				w http.ResponseWriter,
				r *http.Request,
			) {
				defer func() {
					if v := recover(); v != nil {
						if h != nil {
							h.Recover(v)
						}

						http.Error(
							w,
							constant.InternalError,
							http.StatusInternalServerError,
						)
					}
				}()
				n.ServeHTTP(w, r)
			},
		)
	}
}
