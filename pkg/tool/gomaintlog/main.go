package gomaintlog

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/maintenance_log"
	"github.com/funtimecoding/go-library/pkg/monitor"
)

func Main(
	version string,
	gitHash string,
	buildDate string,
) {
	monitor.ParseBind(version, gitHash, buildDate)
	c := maintenance_log.NewEnvironment()
	fmt.Printf("Entries: %s\n", c.Entries())
}
