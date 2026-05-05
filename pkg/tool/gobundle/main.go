package gobundle

import (
	"github.com/funtimecoding/go-library/pkg/argument"
	"github.com/funtimecoding/go-library/pkg/errors/sentry/reporter"
	"github.com/funtimecoding/go-library/pkg/monitor"
	"github.com/funtimecoding/go-library/pkg/tool/gobundle/constant"
	"github.com/funtimecoding/go-library/pkg/tool/gobundle/option"
)

func Main(
	version string,
	gitHash string,
	buildDate string,
) {
	r := reporter.New(constant.Name, version).Start()
	defer func() { r.RecoverFlush(recover()) }()
	monitor.ParseBind(version, gitHash, buildDate)
	o := option.New()
	o.Name = argument.RequiredPositional(0, "NAME")
	o.Path = argument.RequiredPositional(1, "PATH")
	o.Executable = argument.RequiredPositional(2, "EXECUTABLE")
	o.Icon = argument.RequiredPositional(3, "ICON")
	o.Vendor = argument.RequiredPositional(4, "VENDOR")
	o.BundleVersion = argument.RequiredPositional(5, "VERSION")
	Run(o)
}
