package goraidparsed

import (
	"context"
	"github.com/funtimecoding/go-library/pkg/lifecycle"
	"github.com/funtimecoding/go-library/pkg/log/logger"
	"github.com/funtimecoding/go-library/pkg/tool/goraidparsed/option"
	"github.com/funtimecoding/go-library/pkg/tool/goraidparsed/route"
	generated "github.com/funtimecoding/go-library/pkg/tool/goraidparsed/server"
	"github.com/funtimecoding/go-library/pkg/web"
	"github.com/getsentry/sentry-go"
	"net/http"
)

func Run(
	o *option.Parser,
	h *sentry.Hub,
) {
	l := logger.New(context.Background())
	lifecycle.New(
		l,
		lifecycle.WithServerMiddleware(
			":8081",
			func(m *http.ServeMux) {
				generated.HandlerFromMux(
					route.New(
						o.ParserPath,
						o.TemplatePath,
						o.OutputPath,
						l,
						h,
					),
					m,
				)
			},
			web.RecoveryMiddleware(h),
		),
	).RunUntilSignal()
}
