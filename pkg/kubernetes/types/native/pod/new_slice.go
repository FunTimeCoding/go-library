package pod

import core "k8s.io/api/core/v1"

func NewSlice(
	v []core.Pod,
	cluster string,
) []*Pod {
	var result []*Pod

	for _, p := range v {
		result = append(result, New(&p, cluster))
	}

	return result
}
