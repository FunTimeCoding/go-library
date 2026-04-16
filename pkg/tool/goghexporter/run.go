package goghexporter

import (
	"github.com/funtimecoding/go-library/pkg/github"
	"github.com/funtimecoding/go-library/pkg/lifecycle"
	"github.com/funtimecoding/go-library/pkg/metric"
	"github.com/funtimecoding/go-library/pkg/tool/goghexporter/option"
	"github.com/funtimecoding/go-library/pkg/tool/goghexporter/poller"
)

func Run(o *option.Exporter) {
	m := metric.New(0, o.Verbose, nil)
	p := poller.New(
		github.NewEnvironment(),
		o.Owner,
		o.Interval,
		m.Registry(),
	)
	l := lifecycle.New(
		lifecycle.WithVerbose(o.Verbose),
		lifecycle.WithWorker(m),
		lifecycle.WithWorker(p),
	)
	l.RunUntilSignal()
}
