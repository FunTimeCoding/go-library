package job

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/kubernetes/client"
	"github.com/funtimecoding/go-library/pkg/kubernetes/constant"
)

func renovate(
	k *client.Client,
	wait bool,
) {
	deletePreviousManualJob(k, constant.RenovateNamespace)
	j := k.CreateJobFromCron(
		constant.RenovateNamespace,
		constant.RenovateCron,
		constant.ManualCron,
	)
	fmt.Printf("Job: %s\n", j.Name)

	if wait {
		waitForFinish(k, constant.RenovateNamespace, j.Name)
		printJobLog(k, constant.RenovateNamespace, j.Name)
	}
}
