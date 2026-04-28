package goraidparsed

import (
	"github.com/funtimecoding/go-library/pkg/errors/sentry/constant"
	"github.com/funtimecoding/go-library/pkg/errors/sentry/reporter"
	"github.com/funtimecoding/go-library/pkg/monitor"
	"github.com/funtimecoding/go-library/pkg/system/environment"
	"github.com/funtimecoding/go-library/pkg/tool/goraidparsed/option"
	"github.com/getsentry/sentry-go"
)

func Main(
	version string,
	gitHash string,
	buildDate string,
) {
	var h *sentry.Hub

	if c := environment.Optional(constant.LocatorEnvironment); c != "" {
		r := reporter.New("goraidparsed", c, "", version)
		r.Start()
		defer func() { r.RecoverFlush(recover()) }()
		h = r.Hub()
	}

	monitor.ParseBind(version, gitHash, buildDate)
	o := option.New()
	o.ParserPath = "/opt/parser"
	o.TemplatePath = "/opt/template/TW5_Top_Stat_Parse.html"
	o.OutputPath = "/srv/gw2-report"
	Run(o, h)
}
