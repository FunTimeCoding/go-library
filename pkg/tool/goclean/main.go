package goclean

import (
	"github.com/funtimecoding/go-library/pkg/argument"
	"github.com/funtimecoding/go-library/pkg/errors/sentry/reporter"
	gitlab "github.com/funtimecoding/go-library/pkg/gitlab/constant"
	"github.com/funtimecoding/go-library/pkg/system/environment"
	"github.com/funtimecoding/go-library/pkg/tool/goclean/clean"
	"github.com/funtimecoding/go-library/pkg/tool/goclean/clean/option"
	"github.com/funtimecoding/go-library/pkg/tool/goclean/constant"
)

func Main(
	version string,
	gitHash string,
	buildDate string,
) {
	r := reporter.New(constant.Identity.Name(), version).Start()
	defer func() { r.RecoverFlush(recover()) }()
	a := argument.NewInstance(constant.Identity)
	a.Boolean(argument.Verbose, false, "Verbose output")
	a.Parse(version, gitHash, buildDate)
	o := option.New()
	o.GitLabHost = environment.Required(gitlab.HostEnvironment)
	o.Verbose = a.GetBoolean(argument.Verbose)
	clean.Run(o)
}
