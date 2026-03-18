package goalertlogd

import (
	"context"
	generative "github.com/funtimecoding/go-library/pkg/generative/model_context/server"
	"github.com/funtimecoding/go-library/pkg/lifecycle"
	"github.com/funtimecoding/go-library/pkg/log/logger"
	"github.com/funtimecoding/go-library/pkg/metric"
	"github.com/funtimecoding/go-library/pkg/prometheus/alertmanager"
	"github.com/funtimecoding/go-library/pkg/tool/goalertlogd/model_context"
	"github.com/funtimecoding/go-library/pkg/tool/goalertlogd/option"
	"github.com/funtimecoding/go-library/pkg/tool/goalertlogd/poller"
	"github.com/funtimecoding/go-library/pkg/tool/goalertlogd/route"
	generated "github.com/funtimecoding/go-library/pkg/tool/goalertlogd/server"
	"github.com/funtimecoding/go-library/pkg/tool/goalertlogd/store"
	"github.com/funtimecoding/go-library/pkg/tool/goalertlogd/web"
	webConstant "github.com/funtimecoding/go-library/pkg/web/constant"
	"net/http"
	"time"
)

func Run(o *option.Log) {
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
	l := lifecycle.New(
		lifecycle.WithWorker(e),
		lifecycle.WithWorker(p),
		lifecycle.WithServer(
			webConstant.Listen,
			func(m *http.ServeMux) {
				generated.HandlerFromMux(route.New(s, p), m)
				generative.New(model_context.New(s, p).Nested()).Setup(m)
				web.NewServer(s, p).Mount(m)
			},
		),
	)
	g.Structured("starting")
	l.RunUntilSignal()
}
