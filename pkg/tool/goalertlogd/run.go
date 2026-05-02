package goalertlogd

import (
	"context"
	"github.com/funtimecoding/go-library/pkg/face"
	"github.com/funtimecoding/go-library/pkg/lifecycle"
	"github.com/funtimecoding/go-library/pkg/log/logger"
	"github.com/funtimecoding/go-library/pkg/metric"
	"github.com/funtimecoding/go-library/pkg/prometheus/alertmanager"
	generated "github.com/funtimecoding/go-library/pkg/tool/goalertlogd/generated/server"
	"github.com/funtimecoding/go-library/pkg/tool/goalertlogd/model_context"
	"github.com/funtimecoding/go-library/pkg/tool/goalertlogd/option"
	"github.com/funtimecoding/go-library/pkg/tool/goalertlogd/server"
	"github.com/funtimecoding/go-library/pkg/tool/goalertlogd/store"
	alertWeb "github.com/funtimecoding/go-library/pkg/tool/goalertlogd/web"
	"github.com/funtimecoding/go-library/pkg/tool/goalertlogd/worker"
	"github.com/funtimecoding/go-library/pkg/web"
	"github.com/funtimecoding/go-library/pkg/web/constant"
	"net/http"
	"time"
)

func Run(
	o *option.Log,
	r face.Reporter,
) {
	g := logger.New(context.Background())
	m := metric.New(0, false, nil)
	s := store.New(o.DatabasePath)
	defer s.Close()
	w := worker.New(
		alertmanager.NewEnvironment(),
		s,
		g,
		1*time.Minute,
		30*24*time.Hour,
		m.Registry(),
	)
	lifecycle.New(
		g,
		lifecycle.WithWorker(m),
		lifecycle.WithWorker(w),
		lifecycle.WithServerMiddleware(
			constant.ListenAddress,
			func(m *http.ServeMux) {
				generated.HandlerFromMux(server.New(s, w), m)
				model_context.New(s, w, r).Mount(m)
				alertWeb.New(s, w).Mount(m)
			},
			web.RecoveryMiddleware(r),
		),
	).RunUntilSignal()
}
