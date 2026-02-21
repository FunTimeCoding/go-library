package gojira

import (
	"github.com/funtimecoding/go-library/pkg/argument"
	"github.com/funtimecoding/go-library/pkg/atlassian/jira/check/issue"
	"github.com/funtimecoding/go-library/pkg/atlassian/jira/check/issue/option"
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
	o.Notation = viper.GetBool(argument.Notation)
	o.Copyable = viper.GetBool(argument.Copyable)
	o.All = viper.GetBool(argument.All)
	issue.Check(o)
}
