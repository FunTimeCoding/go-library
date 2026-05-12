package godebian

import (
	"github.com/funtimecoding/go-library/pkg/argument"
	"github.com/funtimecoding/go-library/pkg/errors/sentry/reporter"
	"github.com/funtimecoding/go-library/pkg/semver"
	"github.com/funtimecoding/go-library/pkg/tool/godebian/constant"
	"github.com/funtimecoding/go-library/pkg/tool/godebian/option"
)

func Main(
	version string,
	gitHash string,
	buildDate string,
) {
	r := reporter.New(constant.Identity.Name(), version).Start()
	defer func() { r.RecoverFlush(recover()) }()
	a := argument.NewInstance(constant.Identity)
	a.String(argument.Executable, "", "Executable to package: go-example")
	a.String(
		argument.Version,
		"",
		"Package semantic version: 1.0.0, v-prefix gets trimmed",
	)
	a.String(maintainerNameArgument, "", "Maintainer name: AN Other")
	a.String(maintainerEmailArgument, "", "Maintainer email: another@example.org")
	a.Boolean(systemdUnitFlag, false, "Create a systemd unit")
	a.Parse(version, gitHash, buildDate)
	o := option.New()
	o.Executable = a.GetString(argument.Executable)
	o.PackageVersion = semver.Trim(a.GetString(argument.Version))
	o.MaintainerName = a.GetString(maintainerNameArgument)
	o.MaintainerEmail = a.GetString(maintainerEmailArgument)
	o.SystemdUnit = a.GetBoolean(systemdUnitFlag)
	Run(o)
}
