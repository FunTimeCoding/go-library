package goansibled

import (
	"github.com/funtimecoding/go-library/pkg/argument"
	"github.com/funtimecoding/go-library/pkg/errors/sentry/reporter"
	"github.com/funtimecoding/go-library/pkg/monitor"
	"github.com/funtimecoding/go-library/pkg/tool/goansibled/constant"
	"github.com/funtimecoding/go-library/pkg/tool/goansibled/option"
	"github.com/spf13/pflag"
)

func Main(
	version string,
	gitHash string,
	buildDate string,
) {
	r := reporter.New(constant.Name, version).Start()
	defer func() { r.RecoverFlush(recover()) }()
	pflag.String(argument.Repository, "", "Git repository URL")
	pflag.String(argument.ClonePath, "", "Local repository path")
	pflag.String(
		argument.AnsiblePath,
		"",
		"Path within repository to ansible root",
	)
	pflag.String(
		argument.Playbook,
		"",
		"Comma-separated playbook paths relative to ansible path",
	)
	monitor.ParseBind(version, gitHash, buildDate)
	o := option.New()
	o.Repository = argument.Required(argument.Repository)
	o.ClonePath = argument.Required(argument.ClonePath)
	o.AnsiblePath = argument.Required(argument.AnsiblePath)
	o.Playbook = argument.Slice(argument.Playbook)
	Run(o, r)
}
