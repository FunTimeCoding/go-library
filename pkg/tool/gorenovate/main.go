package gorenovate

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/kubernetes/client"
	"github.com/funtimecoding/go-library/pkg/kubernetes/constant"
	"github.com/funtimecoding/go-library/pkg/monitor"
)

func Main(
	version string,
	gitHash string,
	buildDate string,
) {
	monitor.ParseBind(version, gitHash, buildDate)
	k := client.NewEnvironment()
	f := constant.Format

	for _, j := range k.CronJobs(constant.RenovateNamespace) {
		fmt.Printf("CronJob: %s\n", j.Format(f))
	}

	for _, j := range k.Jobs(constant.RenovateNamespace) {
		fmt.Printf("Job: %s\n", j.Format(f))
	}
}
