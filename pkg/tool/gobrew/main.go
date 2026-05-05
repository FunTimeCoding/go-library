package gobrew

import (
	"github.com/funtimecoding/go-library/pkg/argument"
	"github.com/funtimecoding/go-library/pkg/errors/sentry/reporter"
	"github.com/funtimecoding/go-library/pkg/monitor"
	"github.com/funtimecoding/go-library/pkg/system/macos/check/brew/outdated"
	"github.com/funtimecoding/go-library/pkg/system/macos/check/brew/outdated/option"
	"github.com/spf13/viper"
)

func Main(
	version string,
	gitHash string,
	buildDate string,
) {
	r := reporter.New("gobrew", version)
	r.Start()
	defer func() { r.RecoverFlush(recover()) }()
	monitor.CopyableArgument()
	monitor.NotationArgument()
	monitor.AllArgument()
	monitor.ParseBind(version, gitHash, buildDate)
	o := option.New()
	o.Copyable = viper.GetBool(argument.Copyable)
	o.Notation = viper.GetBool(argument.Notation)
	o.All = viper.GetBool(argument.All)
	outdated.Check(o)
}
