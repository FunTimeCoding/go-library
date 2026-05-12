package goatlassiand

import (
	"github.com/funtimecoding/go-library/pkg/argument"
	"github.com/funtimecoding/go-library/pkg/errors/sentry/reporter"
	"github.com/funtimecoding/go-library/pkg/tool/goatlassiand/constant"
	"github.com/funtimecoding/go-library/pkg/tool/goatlassiand/option"
	web "github.com/funtimecoding/go-library/pkg/web/constant"
)

func Main(
	version string,
	gitHash string,
	buildDate string,
) {
	r := reporter.New(constant.Identity.Name(), version).Start()
	defer func() { r.RecoverFlush(recover()) }()
	a := argument.NewInstance(constant.Identity)
	a.Integer(argument.Port, web.ListenPort, web.PortUsage)
	a.Parse(version, gitHash, buildDate)
	o := option.New()
	o.Port = a.RequiredInteger(argument.Port)
	o.Version = version
	Run(o, r)
}
