package namespace

import core "k8s.io/api/core/v1"

func New(
	v *core.Namespace,
	cluster string,
) *Namespace {
	return &Namespace{Cluster: cluster, Name: v.Name, Raw: v}
}
