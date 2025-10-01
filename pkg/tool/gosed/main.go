package gosed

import (
	"github.com/funtimecoding/go-library/pkg/argument"
	"github.com/funtimecoding/go-library/pkg/git/constant"
	"github.com/funtimecoding/go-library/pkg/monitor"
	"github.com/funtimecoding/go-library/pkg/tool/common"
	"github.com/funtimecoding/go-library/pkg/tool/gosed/sed"
	"github.com/funtimecoding/go-library/pkg/tool/gosed/sed/option"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

func Main(
	version string,
	gitHash string,
	buildDate string,
) {
	common.Arguments()
	pflag.String(
		argument.Branch,
		constant.MainBranch,
		"Branch to commit to",
	)
	pflag.String(argument.Path, "", "Path in repository")
	var replaces []string
	pflag.StringSliceVar(
		&replaces,
		argument.Replace,
		nil,
		"One or more prefix replaces (Example: 'image: app:=v1.2.3')",
	)
	monitor.ParseBind(version, gitHash, buildDate)
	common.ValidateArguments()
	o := option.New()
	o.Host = viper.GetString(argument.Host)
	o.Token = viper.GetString(argument.Token)
	o.Owner = viper.GetString(argument.Owner)
	o.Repository = viper.GetString(argument.Repository)
	o.Branch = viper.GetString(argument.Branch)
	o.Path = viper.GetString(argument.Path)
	o.Replaces = sed.Parse(replaces)
	o.Message = argument.RequiredPositional(0, "MESSAGE")
	sed.Run(o)
}
