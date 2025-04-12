package client

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

func New(c *rest.Config) *kubernetes.Clientset {
	result, e := kubernetes.NewForConfig(c)
	errors.PanicOnError(e)

	return result
}
