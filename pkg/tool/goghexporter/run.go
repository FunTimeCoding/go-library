package goghexporter

import (
	"context"
	"github.com/funtimecoding/go-library/pkg/github"
	"github.com/funtimecoding/go-library/pkg/lifecycle"
	"github.com/funtimecoding/go-library/pkg/log/logger"
	"github.com/funtimecoding/go-library/pkg/metric"
	"github.com/funtimecoding/go-library/pkg/tool/goghexporter/option"
	"github.com/funtimecoding/go-library/pkg/tool/goghexporter/poller"
	"github.com/getsentry/sentry-go"
)

func Run(
	o *option.Exporter,
	h *sentry.Hub,
) {
	l := logger.New(context.Background())
	m := metric.New(0, o.Verbose, nil)
	lifecycle.New(
		l,
		lifecycle.WithVerbose(o.Verbose),
		lifecycle.WithWorker(m),
		lifecycle.WithWorker(
			poller.New(
				github.NewEnvironment(),
				o.Owner,
				o.Interval,
				m.Registry(),
				l,
				h,
			),
		),
	).RunUntilSignal()
}
