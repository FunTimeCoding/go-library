package gocommit

import (
	"github.com/funtimecoding/go-library/pkg/argument"
	"github.com/funtimecoding/go-library/pkg/tool/common"
	"github.com/funtimecoding/go-library/pkg/tool/gocommit/commit"
	"github.com/funtimecoding/go-library/pkg/tool/gocommit/commit/option"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

func Main() {
	common.Arguments()
	pflag.String(argument.Branch, "main", "Branch to commit to")
	pflag.String(argument.Path, "", "Path in repository")
	pflag.String(
		argument.Template,
		"",
		"Template file for commit",
	)
	var replaces []string
	pflag.StringSliceVar(
		&replaces,
		argument.Replace,
		nil,
		"One or more key-value pairs to replace (Example: FOO=BAR)",
	)
	argument.ParseBind()
	common.ValidateArguments()
	o := option.New()
	o.Host = viper.GetString(argument.Host)
	o.Token = viper.GetString(argument.Token)
	o.Owner = viper.GetString(argument.Owner)
	o.Repository = viper.GetString(argument.Repository)
	o.Branch = viper.GetString(argument.Branch)
	o.Path = viper.GetString(argument.Path)
	o.Template = viper.GetString(argument.Template)
	o.Replace = replaces
	o.Message = argument.RequiredPositional(0, "MESSAGE")
	commit.Run(o)
}
