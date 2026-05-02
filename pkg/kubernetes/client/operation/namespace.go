package operation

import (
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/kubernetes/typed/core/v1"
)

func Namespace(c *kubernetes.Clientset) v1.NamespaceInterface {
	return c.CoreV1().Namespaces()
}
