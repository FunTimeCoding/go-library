package namespace

import "k8s.io/api/core/v1"

type Namespace struct {
	Cluster string
	Name    string
	Raw     *v1.Namespace
}
