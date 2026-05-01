package goalertlogd

import (
	"context"
	"github.com/funtimecoding/go-library/pkg/lifecycle"
	"github.com/funtimecoding/go-library/pkg/log/logger"
	"github.com/funtimecoding/go-library/pkg/metric"
	"github.com/funtimecoding/go-library/pkg/prometheus/alertmanager"
	generated "github.com/funtimecoding/go-library/pkg/tool/goalertlogd/generated/server"
	"github.com/funtimecoding/go-library/pkg/tool/goalertlogd/model_context"
	"github.com/funtimecoding/go-library/pkg/tool/goalertlogd/option"
	"github.com/funtimecoding/go-library/pkg/tool/goalertlogd/poller"
	"github.com/funtimecoding/go-library/pkg/tool/goalertlogd/server"
	"github.com/funtimecoding/go-library/pkg/tool/goalertlogd/store"
	alertWeb "github.com/funtimecoding/go-library/pkg/tool/goalertlogd/web"
	"github.com/funtimecoding/go-library/pkg/web"
	webConstant "github.com/funtimecoding/go-library/pkg/web/constant"
	"github.com/getsentry/sentry-go"
	"net/http"
	"time"
)

func Run(
	o *option.Log,
	h *sentry.Hub,
) {
	g := logger.New(context.Background())
	e := metric.New(0, false, nil)
	s := store.New(o.DatabasePath)
	defer s.Close()
	p := poller.New(
		alertmanager.NewEnvironment(),
		s,
		g,
		1*time.Minute,
		30*24*time.Hour,
		e.Registry(),
	)
	lifecycle.New(
		g,
		lifecycle.WithWorker(e),
		lifecycle.WithWorker(p),
		lifecycle.WithServerMiddleware(
			webConstant.ListenAddress,
			func(m *http.ServeMux) {
				generated.HandlerFromMux(server.New(s, p), m)
				model_context.New(s, p, h).Mount(m)
				alertWeb.New(s, p).Mount(m)
			},
			web.RecoveryMiddleware(h),
		),
	).RunUntilSignal()
}
