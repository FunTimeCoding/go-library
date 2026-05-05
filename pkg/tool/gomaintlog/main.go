package gomaintlog

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/errors/sentry/reporter"
	"github.com/funtimecoding/go-library/pkg/monitor"
	"github.com/funtimecoding/go-library/pkg/tool/gomaintlog/constant"
	"github.com/funtimecoding/go-library/pkg/tool/gomaintlogd/client"
)

func Main(
	version string,
	gitHash string,
	buildDate string,
) {
	r := reporter.New(constant.Name, version).Start()
	defer func() { r.RecoverFlush(recover()) }()
	monitor.ParseBind(version, gitHash, buildDate)
	c := client.NewEnvironment()
	fmt.Printf("Entries: %s\n", c.Entries())
}
