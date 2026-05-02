package goghexporter

import (
	"context"
	"github.com/funtimecoding/go-library/pkg/face"
	"github.com/funtimecoding/go-library/pkg/github"
	"github.com/funtimecoding/go-library/pkg/lifecycle"
	"github.com/funtimecoding/go-library/pkg/log/logger"
	"github.com/funtimecoding/go-library/pkg/metric"
	"github.com/funtimecoding/go-library/pkg/tool/goghexporter/option"
	"github.com/funtimecoding/go-library/pkg/tool/goghexporter/worker"
)

func Run(
	o *option.Exporter,
	r face.Reporter,
) {
	l := logger.New(context.Background())
	m := metric.New(0, o.Verbose, nil)
	lifecycle.New(
		l,
		lifecycle.WithVerbose(o.Verbose),
		lifecycle.WithWorker(m),
		lifecycle.WithWorker(
			worker.New(
				github.NewEnvironment(),
				o.Owner,
				o.Interval,
				m.Registry(),
				l,
				r,
			),
		),
	).RunUntilSignal()
}
