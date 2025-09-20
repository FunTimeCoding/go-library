package operation

import (
	"k8s.io/client-go/kubernetes"
	batch "k8s.io/client-go/kubernetes/typed/batch/v1"
)

func Job(
	c *kubernetes.Clientset,
	namespace string,
) batch.JobInterface {
	return c.BatchV1().Jobs(namespace)
}
