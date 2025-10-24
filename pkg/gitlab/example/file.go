package example

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/argument"
	"github.com/funtimecoding/go-library/pkg/gitlab"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

func File() {
	pflag.String(argument.Owner, "", "group or user")
	pflag.String(argument.Repository, "", "repository name")
	pflag.String(argument.Branch, "", "branch name")
	pflag.String(argument.File, "", "file path")
	argument.ParseBind()
	owner := viper.GetString(argument.Owner)
	repository := viper.GetString(argument.Repository)
	branch := viper.GetString(argument.Branch)
	file := viper.GetString(argument.File)
	g := gitlab.NewEnvironment()
	p := g.ProjectByName(owner, repository)
	f := g.File(p.Identifier, branch, file)
	fmt.Printf("File: %+v\n", f)
}
