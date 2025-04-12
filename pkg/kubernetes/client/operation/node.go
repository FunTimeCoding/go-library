package operation

import (
	"k8s.io/client-go/kubernetes"
	core "k8s.io/client-go/kubernetes/typed/core/v1"
)

func Node(c *kubernetes.Clientset) core.NodeInterface {
	return c.CoreV1().Nodes()
}
