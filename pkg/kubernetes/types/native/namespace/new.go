package namespace

import "k8s.io/api/core/v1"

func New(
	v *v1.Namespace,
	cluster string,
) *Namespace {
	return &Namespace{Cluster: cluster, Name: v.Name, Raw: v}
}
