package god8y

import (
	"github.com/funtimecoding/go-library/pkg/errors/sentry/reporter"
	"github.com/funtimecoding/go-library/pkg/monitor"
	"github.com/funtimecoding/go-library/pkg/text/dictionary/check/duplicate"
	"github.com/funtimecoding/go-library/pkg/text/dictionary/check/missing"
	"github.com/funtimecoding/go-library/pkg/text/dictionary/check/order"
	"github.com/funtimecoding/go-library/pkg/tool/god8y/constant"
	"os"
)

func Main(
	version string,
	gitHash string,
	buildDate string,
) {
	r := reporter.New(constant.Name, version).Start()
	defer func() { r.RecoverFlush(recover()) }()

	if len(os.Args) >= 3 && os.Args[1] == "merge" {
		merge(os.Args[2:])

		return
	}

	monitor.ParseBind(version, gitHash, buildDate)
	order.Run()
	missing.Run()
	duplicate.Run()
}
