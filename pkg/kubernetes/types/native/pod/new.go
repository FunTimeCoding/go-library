package pod

import core "k8s.io/api/core/v1"

func New(
	v *core.Pod,
	cluster string,
) *Pod {
	return &Pod{Cluster: cluster, Name: v.Name, Raw: v}
}
