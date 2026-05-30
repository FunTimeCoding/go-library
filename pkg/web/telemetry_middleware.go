package web

import (
	"context"
	"github.com/funtimecoding/go-library/pkg/face"
	"github.com/funtimecoding/go-library/pkg/telemetry/constant"
	"github.com/funtimecoding/go-library/pkg/telemetry/record"
	"github.com/oapi-codegen/runtime/strictmiddleware/nethttp"
	"net/http"
)

func TelemetryMiddleware(c face.Recorder) nethttp.StrictHTTPMiddlewareFunc {
	return func(
		f nethttp.StrictHTTPHandlerFunc,
		operation string,
	) nethttp.StrictHTTPHandlerFunc {
		return func(
			x context.Context,
			w http.ResponseWriter,
			r *http.Request,
			request any,
		) (any, error) {
			response, e := f(x, w, r, request)
			outcome := constant.Success

			if e != nil {
				outcome = constant.Error
			}

			c.Record(
				record.NewBaseline(
					operation,
					constant.WebService,
					"user",
					outcome,
				),
			)

			return response, e
		}
	}
}
