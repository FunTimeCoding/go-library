package god8y

import (
	"github.com/funtimecoding/go-library/pkg/errors/sentry/constant"
	"github.com/funtimecoding/go-library/pkg/errors/sentry/reporter"
	"github.com/funtimecoding/go-library/pkg/monitor"
	"github.com/funtimecoding/go-library/pkg/system/environment"
	"github.com/funtimecoding/go-library/pkg/text/dictionary/check/duplicate"
	"github.com/funtimecoding/go-library/pkg/text/dictionary/check/missing"
	"github.com/funtimecoding/go-library/pkg/text/dictionary/check/order"
	"os"
)

func Main(
	version string,
	gitHash string,
	buildDate string,
) {
	if c := environment.Optional(constant.LocatorEnvironment); c != "" {
		r := reporter.New("god8y", c, "", version)
		r.Start()
		defer func() { r.RecoverFlush(recover()) }()
	}

	if len(os.Args) >= 3 && os.Args[1] == "merge" {
		merge(os.Args[2:])

		return
	}

	monitor.ParseBind(version, gitHash, buildDate)
	order.Run()
	missing.Run()
	duplicate.Run()
}
