package goalertlog

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/alert_log"
	"github.com/funtimecoding/go-library/pkg/monitor"
)

func Main(
	version string,
	gitHash string,
	buildDate string,
) {
	monitor.ParseBind(version, gitHash, buildDate)
	l := alert_log.NewEnvironment()
	fmt.Printf("Alerts: %s\n", l.Alerts())
}
