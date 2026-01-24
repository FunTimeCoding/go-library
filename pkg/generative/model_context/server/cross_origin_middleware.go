package server

import (
	"github.com/funtimecoding/go-library/pkg/strings/join"
	web "github.com/funtimecoding/go-library/pkg/web/constant"
	"net/http"
)

func crossOriginMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(
		func(
			w http.ResponseWriter,
			r *http.Request,
		) {
			w.Header().Set(web.AccessOrigin, web.OriginAll)
			w.Header().Set(
				web.AccessHeader,
				join.CommaSpace(
					[]string{
						web.ContentType,
						web.Authorization,
						web.SessionIdentifier,
						web.ProtocolVersion,
					},
				),
			)
			w.Header().Set(
				web.AccessMethod,
				join.CommaSpace(
					[]string{
						http.MethodGet,
						http.MethodPost,
						http.MethodDelete,
						http.MethodOptions,
					},
				),
			)
			w.Header().Set(
				web.AccessExpose,
				join.CommaSpace(
					[]string{web.SessionIdentifier, web.ProtocolVersion},
				),
			)

			if r.Method == http.MethodOptions {
				w.WriteHeader(http.StatusOK)

				return
			}

			next.ServeHTTP(w, r)
		},
	)
}
