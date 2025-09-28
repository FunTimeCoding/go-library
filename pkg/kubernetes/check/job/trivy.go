package job

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/kubernetes/client"
	"github.com/funtimecoding/go-library/pkg/kubernetes/constant"
)

func trivy(
	k *client.Client,
	wait bool,
) {
	deletePreviousManualJob(k, constant.TrivyNamespace)
	j := k.CreateJobFromCron(
		constant.TrivyNamespace,
		constant.TrivyCron,
		constant.ManualCron,
	)
	fmt.Printf("Job: %s\n", j.Name)

	if wait {
		waitForFinish(k, constant.TrivyNamespace, j.Name)
		printJobLog(k, constant.TrivyNamespace, j.Name)
	}
}
