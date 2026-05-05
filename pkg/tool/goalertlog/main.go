package goalertlog

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/errors/sentry/reporter"
	"github.com/funtimecoding/go-library/pkg/monitor"
	"github.com/funtimecoding/go-library/pkg/tool/goalertlogd/client"
)

func Main(
	version string,
	gitHash string,
	buildDate string,
) {
	r := reporter.New("goalertlog", version)
	r.Start()
	defer func() { r.RecoverFlush(recover()) }()
	monitor.ParseBind(version, gitHash, buildDate)
	l := client.NewEnvironment()
	fmt.Printf("Alerts: %s\n", l.Alerts())
}
