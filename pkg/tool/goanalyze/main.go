package goanalyze

import (
	"github.com/funtimecoding/go-library/pkg/argument"
	"github.com/funtimecoding/go-library/pkg/errors/sentry/reporter"
	"github.com/funtimecoding/go-library/pkg/tool/goanalyze/constant"
	"github.com/funtimecoding/go-library/pkg/tool/goanalyze/option"
)

func Main(
	version string,
	gitHash string,
	buildDate string,
) {
	r := reporter.New(constant.Identity.Name(), version).Start()
	defer func() { r.RecoverFlush(recover()) }()
	a := argument.NewInstance(constant.Identity)
	a.Boolean("summary", false, "One line per file")
	a.Parse(version, gitHash, buildDate)
	o := option.New()
	o.Summary = a.GetBoolean("summary")
	o.Patterns = a.Positionals()
	Run(o)
}
