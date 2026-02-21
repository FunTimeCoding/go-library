package gomonitor

import (
	"github.com/funtimecoding/go-library/pkg/argument"
	"github.com/funtimecoding/go-library/pkg/bubbletea"
	monitorModel "github.com/funtimecoding/go-library/pkg/bubbletea/model/monitor"
	"github.com/funtimecoding/go-library/pkg/monitor"
	"github.com/funtimecoding/go-library/pkg/monitor/check/collect"
	"github.com/funtimecoding/go-library/pkg/tool/gomonitor/option"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

func Main(
	version string,
	gitHash string,
	buildDate string,
) {
	pflag.Bool(argument.Connect, false, "Connect to the server")
	pflag.Bool(argument.Once, false, "Run once and exit")
	pflag.Bool(
		argument.DryRun,
		false,
		"Print sources without executing",
	)
	pflag.Bool(argument.Parallel, false, "Run checks in parallel")
	monitor.ParseBind(version, gitHash, buildDate)
	o := option.New()
	o.Once = viper.GetBool(argument.Once)
	o.Connect = viper.GetBool(argument.Connect)
	o.DryRun = viper.GetBool(argument.DryRun)
	o.Parallel = viper.GetBool(argument.Parallel)

	if o.Once {
		collect.Check(o.DryRun, o.Parallel)

		return
	}

	bubbletea.Run(monitorModel.New(o.Connect), true)
}
