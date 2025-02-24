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
	common.ValidateArguments()

	c := gitlab.New(
		viper.GetString(argument.Host),
		viper.GetString(argument.Token),
		[]int{},
	)
	project := common.FindProjectOrExit(
		c,
		viper.GetString(argument.Owner),
		viper.GetString(argument.Repository),
	)
	branch := argument.RequiredStringFlag(argument.Branch, 1)
	path := argument.RequiredStringFlag(argument.Path, 1)
	c.Commit(
		project.Identifier,
		branch,
		argument.RequiredStringFlag(argument.Message, 1),
		path,
		strings.ReplaceAllSlice(
			system.ReadFile(
				argument.RequiredStringFlag(argument.Template, 1),
			),
			replaces,
		),
		c.FileExists(project, branch, path),
	)
}
