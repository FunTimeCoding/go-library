package pod

import "k8s.io/api/core/v1"

func NewSlice(
	v []v1.Pod,
	cluster string,
) []*Pod {
	var result []*Pod

	for _, p := range v {
		result = append(result, New(&p, cluster))
	}

	return result
}
