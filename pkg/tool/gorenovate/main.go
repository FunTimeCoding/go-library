package gorenovate

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/errors/sentry/reporter"
	"github.com/funtimecoding/go-library/pkg/kubernetes/client"
	"github.com/funtimecoding/go-library/pkg/kubernetes/constant"
	"github.com/funtimecoding/go-library/pkg/monitor"
	"os"
)

func Main(
	version string,
	gitHash string,
	buildDate string,
) {
	r := reporter.New("gorenovate", version)
	r.Start()
	defer func() { r.RecoverFlush(recover()) }()
	monitor.ParseBind(version, gitHash, buildDate)
	var missing []string

	if c := parseConfiguration(); c != nil {
		missing = missingFiles(c)

		for _, f := range missing {
			fmt.Printf("matchFiles not found: %s\n", f)
		}
	}

	k := client.NewEnvironment()
	f := constant.Format

	for _, j := range k.CronJobs(constant.RenovateNamespace) {
		fmt.Printf("CronJob: %s\n", j.Format(f))
	}

	for _, j := range k.Jobs(constant.RenovateNamespace) {
		fmt.Printf("Job: %s\n", j.Format(f))
	}

	if len(missing) > 0 {
		os.Exit(1)
	}
}
