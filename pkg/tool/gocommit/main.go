package gocommit

import (
	"github.com/funtimecoding/go-library/pkg/argument"
	"github.com/funtimecoding/go-library/pkg/gitlab"
	"github.com/funtimecoding/go-library/pkg/strings"
	"github.com/funtimecoding/go-library/pkg/system"
	"github.com/funtimecoding/go-library/pkg/tool/common"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

func Main() {
	common.Arguments()
	pflag.String(argument.Branch, "main", "Branch to commit to")
	pflag.String(argument.Message, "", "Commit message")
	pflag.String(argument.Path, "", "Path in repository")
	pflag.String(argument.Template, "", "Template file for commit")
	var replaces []string
	pflag.StringSliceVar(
		&replaces,
		argument.Replace,
		nil,
		"One or more key-value pairs to replace (Example: FOO=BAR)",
	)
	argument.ParseBind()
	common.ValidateArguments()

	c := gitlab.New(
		viper.GetString(argument.Host),
		viper.GetString(argument.Token),
	)
	project := common.FindProjectOrExit(
		c,
		viper.GetString(argument.Owner),
		viper.GetString(argument.Repository),
	)
	branch := argument.RequiredStringFlag(argument.Branch)
	path := argument.RequiredStringFlag(argument.Path)
	c.Commit(
		project.Identifier,
		branch,
		argument.RequiredStringFlag(argument.Message),
		path,
		strings.ReplaceAllSlice(
			system.ReadFile(argument.RequiredStringFlag(argument.Template)),
			replaces,
		),
		c.FileExists(project, branch, path),
	)
}
