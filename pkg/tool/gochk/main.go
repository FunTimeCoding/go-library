package gochk

import (
	"github.com/funtimecoding/go-library/pkg/argument"
	"github.com/funtimecoding/go-library/pkg/errors/sentry/reporter"
	"github.com/funtimecoding/go-library/pkg/tool/gochk/check"
	"github.com/funtimecoding/go-library/pkg/tool/gochk/constant"
)

func Main(
	version string,
	gitHash string,
	buildDate string,
) {
	r := reporter.New(constant.Identity.Name(), version).Start()
	defer func() { r.RecoverFlush(recover()) }()
	a := argument.NewInstance(constant.Identity)
	a.String(
		argument.Port,
		"",
		"Port, multiple values separated by comma",
	)
	a.Parse(version, gitHash, buildDate)
	check.Check()
}
