package gomonitor

import (
	"github.com/funtimecoding/go-library/pkg/argument"
	"github.com/funtimecoding/go-library/pkg/bubbletea"
	"github.com/funtimecoding/go-library/pkg/bubbletea/model/monitor"
	"github.com/funtimecoding/go-library/pkg/monitor/check/collect"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

func Main() {
	pflag.Bool(argument.Connect, false, "Connect to the server")
	pflag.Bool(argument.Once, false, "Run once and exit")
	argument.ParseBind()

	if viper.GetBool(argument.Once) {
		collect.Check()

		return
	}

	bubbletea.Run(monitor.New(viper.GetBool(argument.Connect)), true)
}
