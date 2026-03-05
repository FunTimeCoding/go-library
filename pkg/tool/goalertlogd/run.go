package goalertlogd

import (
	"context"
	"github.com/funtimecoding/go-library/pkg/errors/sentry/constant"
	"github.com/funtimecoding/go-library/pkg/errors/sentry/reporter"
	"github.com/funtimecoding/go-library/pkg/lifecycle"
	"github.com/funtimecoding/go-library/pkg/log/logger"
	"github.com/funtimecoding/go-library/pkg/prometheus/alertmanager"
	"github.com/funtimecoding/go-library/pkg/system"
	"github.com/funtimecoding/go-library/pkg/system/environment"
	"github.com/funtimecoding/go-library/pkg/tool/goalertlogd/api"
	"github.com/funtimecoding/go-library/pkg/tool/goalertlogd/option"
	"github.com/funtimecoding/go-library/pkg/tool/goalertlogd/poller"
	"github.com/funtimecoding/go-library/pkg/tool/goalertlogd/route"
	"github.com/funtimecoding/go-library/pkg/tool/goalertlogd/store"
	webConstant "github.com/funtimecoding/go-library/pkg/web/constant"
	"net/http"
	"time"
)

func Run(o *option.Log) {
	g := logger.New(context.Background())
	locator := environment.Optional(constant.LocatorEnvironment)

	if locator != "" {
		r := reporter.New(
			"goalertlog",
			locator,
			"",
			o.Version,
		)
		r.Start()
		defer func() { r.RecoverFlush(recover()) }()
	}

	s := store.New(o.DatabasePath)
	defer s.Close()
	p := poller.New(
		alertmanager.NewEnvironment(),
		s,
		g,
		1*time.Minute,
		30*24*time.Hour,
	)
	l := lifecycle.New(
		lifecycle.WithWorker(p),
		lifecycle.WithServer(
			webConstant.Listen,
			func(m *http.ServeMux) {
				api.HandlerFromMux(route.New(s, p), m)
			},
		),
	)
	g.Structured("starting")
	l.Run()
	system.KillSignalBlock()
	l.Stop()
}
