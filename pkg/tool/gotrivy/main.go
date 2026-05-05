package gotrivy

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/errors/sentry/reporter"
	"github.com/funtimecoding/go-library/pkg/kubernetes/client"
	kubernetes "github.com/funtimecoding/go-library/pkg/kubernetes/constant"
	"github.com/funtimecoding/go-library/pkg/monitor"
	"github.com/funtimecoding/go-library/pkg/tool/gotrivy/constant"
)

func Main(
	version string,
	gitHash string,
	buildDate string,
) {
	r := reporter.New(constant.Name, version).Start()
	defer func() { r.RecoverFlush(recover()) }()
	monitor.ParseBind(version, gitHash, buildDate)
	k := client.NewEnvironment()
	f := kubernetes.Format

	for _, j := range k.CronJobs(kubernetes.TrivyNamespace) {
		fmt.Printf("CronJob: %s\n", j.Format(f))
	}

	for _, j := range k.Jobs(kubernetes.TrivyNamespace) {
		fmt.Printf("Job: %s\n", j.Format(f))
	}
}
