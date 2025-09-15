package example

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/kubernetes/client"
)

const (
	TrivyNamespace    = "trivy"
	TrivyCron         = "trivy"
	RenovateNamespace = "renovate"
	RenovateCron      = "renovate"
	ManualCron        = "manual"
)

func Job() {
	k := client.NewEnvironment()

	trivy(k)

	if false {
		renovate(k)
	}
}

func renovate(k *client.Client) {
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

func trivy(k *client.Client) {
	for _, j := range k.CronJobs(TrivyNamespace) {
		fmt.Printf("Delete: %s\n", j.Name)
		k.DeleteJobWatch(TrivyNamespace, j.Name)
	}

	if t := k.CronJob(TrivyNamespace, "missing"); t != nil {
		fmt.Printf("Job: %s\n", t.Name)
	}

	k.DeleteJobWatch(TrivyNamespace, ManualCron)
	j := k.CreateJobFromCron(TrivyNamespace, TrivyCron, ManualCron)
	fmt.Printf("Job: %s\n", j.Name)
}
