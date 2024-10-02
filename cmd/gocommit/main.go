package main

import (
	"github.com/funtimecoding/go-library/pkg/argument"
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/gitlab"
	stringsHelper "github.com/funtimecoding/go-library/pkg/strings"
	"github.com/funtimecoding/go-library/pkg/system"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"strings"
)

func main() {
	pflag.String(argument.Host, "", "GitLab host")
	pflag.String(argument.Token, "", "GitLab token")
	pflag.String(
		argument.Owner,
		"",
		"Owner or namespace of project to commit to",
	)
	pflag.String(
		argument.Repository,
		"",
		"Repository to commit to",
	)
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
	pflag.Parse()
	errors.PanicOnError(viper.BindPFlags(pflag.CommandLine))
	content := system.ReadFile(viper.GetString(argument.Template))

	for k, v := range stringsHelper.ToMap(replaces, "=") {
		content = strings.ReplaceAll(content, k, v)
	}

	c := gitlab.New(
		viper.GetString(argument.Host),
		viper.GetString(argument.Token),
		[]int{},
	)
	owner := viper.GetString(argument.Owner)
	repository := viper.GetString(argument.Repository)
	p := c.ProjectByName(owner, repository)

	if p == nil {
		system.Exitf(
			1,
			"repository not found: %s/%s\n",
			owner,
			repository,
		)

		return
	}

	f := c.File(p.Identifier, "main", viper.GetString(argument.Path))
	var update bool

	if f != nil {
		update = true
	}

	c.Commit(
		p.Identifier,
		viper.GetString(argument.Branch),
		viper.GetString(argument.Message),
		viper.GetString(argument.Path),
		content,
		update,
	)
}
