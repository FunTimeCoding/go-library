package operation

import (
	"k8s.io/client-go/kubernetes"
	core "k8s.io/client-go/kubernetes/typed/core/v1"
)

func Pod(
	c *kubernetes.Clientset,
	namespace string,
) core.PodInterface {
	return c.CoreV1().Pods(namespace)
}
