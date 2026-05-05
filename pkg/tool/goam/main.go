package goam

import (
	"github.com/funtimecoding/go-library/pkg/argument"
	"github.com/funtimecoding/go-library/pkg/errors/sentry/reporter"
	"github.com/funtimecoding/go-library/pkg/monitor"
	"github.com/funtimecoding/go-library/pkg/prometheus/amtool"
)

func Main(
	version string,
	gitHash string,
	buildDate string,
) {
	r := reporter.New("goam", version)
	r.Start()
	defer func() { r.RecoverFlush(recover()) }()
	monitor.ParseBind(version, gitHash, buildDate)
	amtool.Run(argument.Positional(0))
}
