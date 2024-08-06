package main

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/gitlab"
	stringsHelper "github.com/funtimecoding/go-library/pkg/strings"
	"github.com/funtimecoding/go-library/pkg/system"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"strings"
)

const HostArgument = "host"
const TokenArgument = "token"
const OwnerArgument = "owner"
const RepositoryArgument = "repository"
const BranchArgument = "branch"
const MessageArgument = "message"
const PathArgument = "path"
const TemplateArgument = "template"
const ReplaceArgument = "replace"

func main() {
	pflag.String(HostArgument, "", "GitLab host")
	pflag.String(TokenArgument, "", "GitLab token")
	pflag.String(
		OwnerArgument,
		"",
		"Owner or namespace of project to commit to",
	)
	pflag.String(
		RepositoryArgument,
		"",
		"Repository to commit to",
	)
	pflag.String(BranchArgument, "main", "Branch to commit to")
	pflag.String(MessageArgument, "", "Commit message")
	pflag.String(PathArgument, "", "Path in repository")
	pflag.String(TemplateArgument, "", "Template file for commit")
	var replaces []string
	pflag.StringSliceVar(
		&replaces,
		ReplaceArgument,
		nil,
		"One or more key-value pairs to replace (Example: FOO=BAR)",
	)
	pflag.Parse()
	errors.PanicOnError(viper.BindPFlags(pflag.CommandLine))
	content := system.ReadFile(viper.GetString(TemplateArgument))

	for k, v := range stringsHelper.ToMap(replaces, "=") {
		content = strings.ReplaceAll(content, k, v)
	}

	c := gitlab.New(
		viper.GetString(HostArgument),
		viper.GetString(TokenArgument),
	)
	owner := viper.GetString(OwnerArgument)
	repository := viper.GetString(RepositoryArgument)
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

	f := c.File(p.Identifier, "main", viper.GetString(PathArgument))
	var update bool

	if f != nil {
		update = true
	}

	c.Commit(
		p.Identifier,
		viper.GetString(BranchArgument),
		viper.GetString(MessageArgument),
		viper.GetString(PathArgument),
		content,
		update,
	)
}
