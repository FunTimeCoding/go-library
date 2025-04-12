package operation

import (
	"k8s.io/client-go/kubernetes"
	core "k8s.io/client-go/kubernetes/typed/core/v1"
)

func Namespace(c *kubernetes.Clientset) core.NamespaceInterface {
	return c.CoreV1().Namespaces()
}
