package goaudit

import (
	"github.com/funtimecoding/go-library/pkg/argument"
	"github.com/funtimecoding/go-library/pkg/errors/sentry/reporter"
	"github.com/funtimecoding/go-library/pkg/tool/goaudit/constant"
	"github.com/funtimecoding/go-library/pkg/tool/goaudit/option"
	"os"
)

func Main(
	version string,
	gitHash string,
	buildDate string,
) {
	r := reporter.New(constant.Identity.Name(), version)
	r.Start()
	defer func() { r.RecoverFlush(recover()) }()
	a := argument.NewInstance(constant.Identity)
	a.Boolean(
		argument.Table,
		false,
		"Print compliance tables instead of lint-style concerns",
	)
	a.Parse(version, gitHash, buildDate)
	roots := a.Positionals()

	if len(roots) == 0 {
		a.PrintUsage()
		os.Exit(1)
	}

	o := option.New()
	o.Roots = roots
	o.Table = a.GetBoolean(argument.Table)
	Run(o)
}
