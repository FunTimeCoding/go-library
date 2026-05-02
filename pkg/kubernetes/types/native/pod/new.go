package pod

import "k8s.io/api/core/v1"

func New(
	v *v1.Pod,
	cluster string,
) *Pod {
	return &Pod{Cluster: cluster, Name: v.Name, Raw: v}
}
