package gomonitor

import (
	"github.com/funtimecoding/go-library/pkg/argument"
	"github.com/funtimecoding/go-library/pkg/bubbletea"
	"github.com/funtimecoding/go-library/pkg/bubbletea/model/monitor"
	"github.com/funtimecoding/go-library/pkg/monitor/check/collect"
	"github.com/funtimecoding/go-library/pkg/tool/gomonitor/option"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

func Main() {
	pflag.Bool(argument.Connect, false, "Connect to the server")
	pflag.Bool(argument.Once, false, "Run once and exit")
	argument.ParseBind()
	o := option.New()
	o.Once = viper.GetBool(argument.Once)
	o.Connect = viper.GetBool(argument.Connect)

	if o.Once {
		collect.Check()

		return
	}

	bubbletea.Run(monitor.New(o.Connect), true)
}
