package example

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/kubernetes/client"
)

const (
	RenovateNamespace = "renovate"
	RenovateCron      = "renovate"
	ManualCron        = "manual"
)

func Job() {
	k := client.NewEnvironment()

	for _, j := range k.CronJobs(RenovateNamespace) {
		fmt.Printf("Delete: %s\n", j.Name)
		k.DeleteJobWatch(RenovateNamespace, j.Name)
	}

	if t := k.CronJob(RenovateNamespace, "missing"); t != nil {
		fmt.Printf("Job: %s\n", t.Name)
	}

	k.DeleteJobWatch(RenovateNamespace, ManualCron)
	j := k.CreateJobFromCron(RenovateNamespace, RenovateCron, ManualCron)
	fmt.Printf("Job: %s\n", j.Name)
}
