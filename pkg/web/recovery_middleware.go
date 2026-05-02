package web

import (
	"github.com/funtimecoding/go-library/pkg/face"
	"github.com/funtimecoding/go-library/pkg/web/constant"
	"net/http"
)

func RecoveryMiddleware(r face.Reporter) func(http.Handler) http.Handler {
	return func(n http.Handler) http.Handler {
		return http.HandlerFunc(
			func(
				w http.ResponseWriter,
				q *http.Request,
			) {
				defer func() {
					if v := recover(); v != nil {
						r.Recover(v)
						http.Error(
							w,
							constant.InternalError,
							http.StatusInternalServerError,
						)
					}
				}()
				n.ServeHTTP(w, q)
			},
		)
	}
}
