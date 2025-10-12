package job

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/kubernetes/client"
	batch "k8s.io/api/batch/v1"
)

func createJob(
	k *client.Client,
	namespace string,
	cron string,
	name string,
) *batch.Job {
	fmt.Printf("Create job %s in %s from %s\n", name, namespace, cron)
	result := k.CreateJobFromCron(namespace, cron, name)
	fmt.Printf("  %s\n", result.Name)

	return result
}
