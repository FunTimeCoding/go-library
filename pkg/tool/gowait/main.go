package gowait

import (
	"github.com/funtimecoding/go-library/pkg/argument"
	"github.com/funtimecoding/go-library/pkg/monitor"
	"github.com/funtimecoding/go-library/pkg/tool/gowait/wait"
	"github.com/funtimecoding/go-library/pkg/tool/gowait/wait/option"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"time"
)

func Main() {
	pflag.String(argument.File, "", "File to wait for")
	pflag.String(argument.Process, "", "Process to wait for")
	pflag.String(argument.Locator, "", "Locator to wait for")
	pflag.String(argument.Contains, "", "String for locator")
	pflag.Duration(argument.Timeout, 3*time.Minute, "")
	monitor.VerboseArgument()
	argument.ParseBind()
	o := option.New()
	o.File = viper.GetString(argument.File)
	o.Process = viper.GetString(argument.Process)
	o.Locator = viper.GetString(argument.Locator)
	o.Timeout = viper.GetDuration(argument.Timeout)
	o.Verbose = viper.GetBool(argument.Verbose)
	wait.Run(o)
}
