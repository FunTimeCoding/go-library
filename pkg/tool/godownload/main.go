package godownload

import (
	"github.com/funtimecoding/go-library/pkg/argument"
	library "github.com/funtimecoding/go-library/pkg/constant"
	"github.com/funtimecoding/go-library/pkg/errors/sentry/reporter"
	"github.com/funtimecoding/go-library/pkg/system/environment"
	"github.com/funtimecoding/go-library/pkg/tool/common"
	"github.com/funtimecoding/go-library/pkg/tool/godownload/constant"
	"github.com/funtimecoding/go-library/pkg/tool/godownload/download"
	"github.com/funtimecoding/go-library/pkg/tool/godownload/download/option"
)

func Main(
	version string,
	gitHash string,
	buildDate string,
) {
	r := reporter.New(constant.Identity.Name(), version).Start()
	defer func() { r.RecoverFlush(recover()) }()
	a := argument.NewInstance(constant.Identity)
	common.Arguments(a)
	a.String(
		argument.PackageVersion,
		library.LatestVersion,
		"Version to download, falls back to latest if not found",
	)
	a.String(
		argument.Output,
		environment.Fallback(
		"OUTPUT",
		download.DefaultOutput,
	),
		"Output directory for executable",
	)
	a.Boolean(argument.Verbose, false, "Verbose output")
	a.Parse(version, gitHash, buildDate)
	common.ValidateArguments(a)
	o := option.New()
	o.Host = a.GetString(argument.Host)
	o.Token = a.GetString(argument.Token)
	o.Owner = a.GetString(argument.Owner)
	o.Repository = a.GetString(argument.Repository)
	o.PackageVersion = a.GetString(argument.PackageVersion)
	o.Output = a.GetString(argument.Output)
	o.Verbose = a.GetBoolean(argument.Verbose)
	o.Package = a.RequiredPositional(0, "PACKAGE")
	download.Run(o)
}
