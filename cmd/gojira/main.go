package main

import (
	"github.com/funtimecoding/go-library/pkg/argument"
	"github.com/funtimecoding/go-library/pkg/atlassian/jira/check/issue"
	"github.com/funtimecoding/go-library/pkg/atlassian/jira/check/issue/option"
	"github.com/funtimecoding/go-library/pkg/monitor"
	"github.com/spf13/viper"
)

func main() {
	monitor.NotationArgument()
	monitor.AllArgument()
	argument.ParseAndBind()
	o := option.New()
	o.All = viper.GetBool(argument.All)
	o.Notation = viper.GetBool(argument.Notation)
	issue.Print(o)
}
