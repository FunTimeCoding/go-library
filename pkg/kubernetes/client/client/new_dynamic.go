package client

import (
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/rest"
)

func NewDynamic(c *rest.Config) (dynamic.Interface, error) {
	return dynamic.NewForConfig(c)
}
