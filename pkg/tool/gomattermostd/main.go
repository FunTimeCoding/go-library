package gomattermostd

import (
	"github.com/funtimecoding/go-library/pkg/argument"
	"github.com/funtimecoding/go-library/pkg/errors/sentry/reporter"
	"github.com/funtimecoding/go-library/pkg/monitor"
	"github.com/funtimecoding/go-library/pkg/tool/gomattermostd/constant"
	"github.com/funtimecoding/go-library/pkg/tool/gomattermostd/option"
	web "github.com/funtimecoding/go-library/pkg/web/constant"
	"github.com/spf13/pflag"
)

func Main(
	version string,
	gitHash string,
	buildDate string,
) {
	r := reporter.New(constant.Name, version).Start()
	defer func() { r.RecoverFlush(recover()) }()
	pflag.Int(argument.Port, web.ListenPort, web.PortUsage)
	monitor.ParseBind(version, gitHash, buildDate)
	o := option.New()
	o.Port = argument.RequiredInteger(argument.Port)
	o.Version = version
	Run(o, r)
}
