package example

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/kubernetes/client"
)

func Job() {
	k := client.NewEnvironment()

	if t := k.CronJob("renovate", "missing"); t != nil {
		fmt.Printf("Job: %s\n", t.Name)
	}

	k.DeleteJobWatch("renovate", "manual")
	j := k.CreateJobFromCron(
		"renovate",
		"renovate",
		"manual",
	)
	fmt.Printf("Job: %s\n", j.Name)
}
