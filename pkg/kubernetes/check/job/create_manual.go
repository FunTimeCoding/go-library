package job

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/kubernetes/client"
	"github.com/funtimecoding/go-library/pkg/kubernetes/constant"
	batch "k8s.io/api/batch/v1"
)

func createManual(
	k *client.Client,
	namespace string,
	cron string,
) *batch.Job {
	name := constant.ManualCron
	fmt.Printf("Create job %s in %s from %s\n", name, namespace, cron)
	result := k.CreateJobFromCron(namespace, cron, name)
	fmt.Printf("  %s\n", result.Name)

	return result
}
