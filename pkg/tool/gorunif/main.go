package gorunif

import (
	"github.com/funtimecoding/go-library/pkg/argument"
	"github.com/funtimecoding/go-library/pkg/errors/sentry/reporter"
	"github.com/funtimecoding/go-library/pkg/monitor"
	"github.com/funtimecoding/go-library/pkg/tool/gorunif/constant"
	"github.com/funtimecoding/go-library/pkg/tool/gorunif/run_if"
	runIfConstant "github.com/funtimecoding/go-library/pkg/tool/gorunif/run_if/constant"
	"github.com/funtimecoding/go-library/pkg/tool/gorunif/run_if/option"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

func Main(
	version string,
	gitHash string,
	buildDate string,
) {
	r := reporter.New(constant.Name, version).Start()
	defer func() { r.RecoverFlush(recover()) }()
	pflag.String(runIfConstant.Base, "HEAD~1", "Base commit")
	pflag.String(runIfConstant.Head, "HEAD", "Head commit")
	monitor.VerboseArgument()
	monitor.ParseBind(version, gitHash, buildDate)
	o := option.New()
	o.Verbose = viper.GetBool(argument.Verbose)
	run_if.Run(o)
}
