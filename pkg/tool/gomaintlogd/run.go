package gomaintlogd

import (
	"context"
	generative "github.com/funtimecoding/go-library/pkg/generative/model_context/server"
	"github.com/funtimecoding/go-library/pkg/lifecycle"
	"github.com/funtimecoding/go-library/pkg/log/logger"
	"github.com/funtimecoding/go-library/pkg/tool/gomaintlogd/model_context"
	"github.com/funtimecoding/go-library/pkg/tool/gomaintlogd/option"
	"github.com/funtimecoding/go-library/pkg/tool/gomaintlogd/route"
	generated "github.com/funtimecoding/go-library/pkg/tool/gomaintlogd/server"
	maintenanceWeb "github.com/funtimecoding/go-library/pkg/tool/gomaintlogd/web"
	"github.com/funtimecoding/go-library/pkg/web"
	"github.com/funtimecoding/go-library/pkg/web/constant"
	"github.com/getsentry/sentry-go"
	"net/http"
)

func Run(
	o *option.Log,
	h *sentry.Hub,
) {
	g := logger.New(context.Background())
	s := newStore(o)
	defer s.Close()
	lifecycle.New(
		g,
		lifecycle.WithServerMiddleware(
			constant.ListenAddress,
			func(m *http.ServeMux) {
				generated.HandlerFromMux(route.New(s), m)
				generative.New(model_context.New(s).Nested()).Setup(m)
				maintenanceWeb.New(s).Mount(m)
			},
			web.RecoveryMiddleware(h),
		),
	).RunUntilSignal()
}
