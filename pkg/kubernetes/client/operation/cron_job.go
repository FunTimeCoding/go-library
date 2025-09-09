package operation

import (
	"k8s.io/client-go/kubernetes"
	batch "k8s.io/client-go/kubernetes/typed/batch/v1"
)

func CronJob(
	c *kubernetes.Clientset,
	namespace string,
) batch.CronJobInterface {
	return c.BatchV1().CronJobs(namespace)
}
