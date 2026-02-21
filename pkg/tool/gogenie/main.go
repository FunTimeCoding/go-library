package gogenie

import (
	"github.com/funtimecoding/go-library/pkg/argument"
	"github.com/funtimecoding/go-library/pkg/atlassian/opsgenie/check/alert"
	"github.com/funtimecoding/go-library/pkg/atlassian/opsgenie/check/alert/option"
	"github.com/funtimecoding/go-library/pkg/monitor"
	"github.com/spf13/viper"
)

func Main(
	version string,
	gitHash string,
	buildDate string,
) {
	monitor.CopyableArgument()
	monitor.NotationArgument()
	monitor.AllArgument()
	monitor.ParseBind(version, gitHash, buildDate)
	o := option.New()
	o.Copyable = viper.GetBool(argument.Copyable)
	o.Notation = viper.GetBool(argument.Notation)
	o.All = viper.GetBool(argument.All)
	alert.Check(o)
}
