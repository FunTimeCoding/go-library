package gomonitor

import (
	"github.com/funtimecoding/go-library/pkg/argument"
	"github.com/funtimecoding/go-library/pkg/bubbletea"
	"github.com/funtimecoding/go-library/pkg/bubbletea/model/monitor"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

func Main() {
	pflag.Bool(
		argument.Connect,
		false,
		"connect to the server",
	)
	argument.ParseBind()
	bubbletea.Run(monitor.New(viper.GetBool(argument.Connect)), true)
}
