package gosentryd

import (
	"context"
	"github.com/funtimecoding/go-library/pkg/errors/sentry"
	"github.com/funtimecoding/go-library/pkg/face"
	"github.com/funtimecoding/go-library/pkg/lifecycle"
	"github.com/funtimecoding/go-library/pkg/log/logger"
	"github.com/funtimecoding/go-library/pkg/tool/gosentryd/model_context"
	"github.com/funtimecoding/go-library/pkg/tool/gosentryd/option"
	"github.com/funtimecoding/go-library/pkg/web"
	"github.com/funtimecoding/go-library/pkg/web/constant"
	"net/http"
)

func Run(
	o *option.Server,
	r face.Reporter,
) {
	lifecycle.New(
		logger.New(context.Background()),
		lifecycle.WithServerMiddleware(
			constant.ListenAddress,
			func(m *http.ServeMux) {
				model_context.New(
					sentry.NewEnvironment(),
					o.Organization,
					r,
					o.Version,
				).Mount(m)
			},
			web.RecoveryMiddleware(r),
		),
	).RunUntilSignal()
}
