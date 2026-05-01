package gomaintlogd

import (
	"context"
	"github.com/funtimecoding/go-library/pkg/lifecycle"
	"github.com/funtimecoding/go-library/pkg/log/logger"
	generated "github.com/funtimecoding/go-library/pkg/tool/gomaintlogd/generated/server"
	"github.com/funtimecoding/go-library/pkg/tool/gomaintlogd/model_context"
	"github.com/funtimecoding/go-library/pkg/tool/gomaintlogd/option"
	"github.com/funtimecoding/go-library/pkg/tool/gomaintlogd/server"
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
				generated.HandlerFromMux(server.New(s), m)
				model_context.New(s, h).Mount(m)
				maintenanceWeb.New(s).Mount(m)
			},
			web.RecoveryMiddleware(h),
		),
	).RunUntilSignal()
}
