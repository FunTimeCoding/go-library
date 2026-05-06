package goansibled

import (
	"github.com/funtimecoding/go-library/pkg/argument"
	"github.com/funtimecoding/go-library/pkg/errors/sentry/reporter"
	"github.com/funtimecoding/go-library/pkg/monitor"
	"github.com/funtimecoding/go-library/pkg/relational/postgres"
	"github.com/funtimecoding/go-library/pkg/system/environment"
	"github.com/funtimecoding/go-library/pkg/tool/goansibled/constant"
	"github.com/funtimecoding/go-library/pkg/tool/goansibled/option"
	web "github.com/funtimecoding/go-library/pkg/web/constant"
	"github.com/spf13/pflag"
)

func Main(
	version string,
	gitHash string,
	buildDate string,
) {
	r := reporter.New(constant.Name, version).Start()
	defer func() { r.RecoverFlush(recover()) }()
	pflag.Int(argument.Port, web.ListenPort, web.PortUsage)
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
	o.Port = argument.RequiredInteger(argument.Port)
	o.Version = version
	o.Repository = argument.Required(argument.Repository)
	o.ClonePath = argument.Required(argument.ClonePath)
	o.AnsiblePath = argument.Required(argument.AnsiblePath)
	o.Playbook = argument.Slice(argument.Playbook)
	o.PostgresLocator = environment.Required(postgres.LocatorEnvironment)
	Run(o, r)
}
