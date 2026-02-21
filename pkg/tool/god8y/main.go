package god8y

import (
	"github.com/funtimecoding/go-library/pkg/monitor"
	"github.com/funtimecoding/go-library/pkg/text/dictionary/check/duplicate"
	"github.com/funtimecoding/go-library/pkg/text/dictionary/check/missing"
	"github.com/funtimecoding/go-library/pkg/text/dictionary/check/order"
)

func Main(
	version string,
	gitHash string,
	buildDate string,
) {
	monitor.ParseBind(version, gitHash, buildDate)
	order.Run()
	missing.Run()
	duplicate.Run()
}
