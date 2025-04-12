package namespace

import core "k8s.io/api/core/v1"

type Namespace struct {
	Cluster string
	Name    string
	Raw     *core.Namespace
}
