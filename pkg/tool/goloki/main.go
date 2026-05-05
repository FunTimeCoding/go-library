package goloki

import (
	"github.com/funtimecoding/go-library/pkg/argument"
	"github.com/funtimecoding/go-library/pkg/errors/sentry/reporter"
	"github.com/funtimecoding/go-library/pkg/monitor"
	"github.com/funtimecoding/go-library/pkg/prometheus/check/loki"
	"github.com/funtimecoding/go-library/pkg/prometheus/check/loki/option"
	lokiConstant "github.com/funtimecoding/go-library/pkg/prometheus/loki/constant"
	"github.com/funtimecoding/go-library/pkg/system/environment"
	"github.com/funtimecoding/go-library/pkg/tool/goloki/constant"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"time"
)

func Main(
	version string,
	gitHash string,
	buildDate string,
) {
	r := reporter.New(constant.Name, version).Start()
	defer func() { r.RecoverFlush(recover()) }()
	monitor.CopyableArgument()
	pflag.Duration(argument.Since, time.Hour, "Time range to query")
	pflag.String(argument.Route, "", "Filter by HTTP route")
	pflag.String(argument.Message, "", "Filter by message field")
	pflag.BoolP(
		argument.Body,
		"b",
		false,
		"Output only the body field",
	)
	pflag.IntP(
		argument.Limit,
		"n",
		10,
		"Maximum number of log entries",
	)
	monitor.ParseBind(version, gitHash, buildDate)
	o := option.New()
	o.Namespace = argument.Positional(0)
	o.Since = viper.GetDuration(argument.Since)
	o.Route = viper.GetString(argument.Route)
	o.Message = viper.GetString(argument.Message)
	o.BodyOnly = viper.GetBool(argument.Body)
	o.Copyable = viper.GetBool(argument.Copyable)
	o.Limit = viper.GetInt(argument.Limit)
	o.Namespaces = environment.Slice(lokiConstant.NamespaceEnvironment)
	o.Exclude = environment.Slice(lokiConstant.ExcludeEnvironment)
	loki.Check(o)
}
