package gogw2

import (
	"github.com/funtimecoding/go-library/pkg/errors/sentry/reporter"
	"github.com/funtimecoding/go-library/pkg/monitor"
	"github.com/funtimecoding/go-library/pkg/tool/gogw2/option"
)

func Main(
	version string,
	gitHash string,
	buildDate string,
) {
	r := reporter.New("gogw2", version)
	r.Start()
	defer func() { r.RecoverFlush(recover()) }()
	monitor.ParseBind(version, gitHash, buildDate)
	o := option.New()
	Run(o)
}
