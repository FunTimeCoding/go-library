package web

import (
	"context"
	"github.com/funtimecoding/go-library/pkg/face"
	"github.com/funtimecoding/go-library/pkg/telemetry/constant"
	"github.com/funtimecoding/go-library/pkg/telemetry/record"
	"net/http"
)

func TelemetryMiddleware(
	c face.Recorder,
) func(
	func(context.Context, http.ResponseWriter, *http.Request, any) (any, error),
	string,
) func(context.Context, http.ResponseWriter, *http.Request, any) (any, error) {
	return func(
		f func(context.Context, http.ResponseWriter, *http.Request, any) (any, error),
		operation string,
	) func(context.Context, http.ResponseWriter, *http.Request, any) (any, error) {
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
