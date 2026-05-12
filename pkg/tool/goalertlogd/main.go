package goalertlogd

import (
	"github.com/funtimecoding/go-library/pkg/argument"
	"github.com/funtimecoding/go-library/pkg/errors/sentry/reporter"
	"github.com/funtimecoding/go-library/pkg/tool/goalertlogd/constant"
	"github.com/funtimecoding/go-library/pkg/tool/goalertlogd/option"
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
		argument.Path,
		"tmp/goalertlog.db",
		"Database path",
	)
	a.Parse(version, gitHash, buildDate)
	o := option.New()
	o.DatabasePath = a.GetString(argument.Path)
	o.Version = version
	Run(o, r)
}
