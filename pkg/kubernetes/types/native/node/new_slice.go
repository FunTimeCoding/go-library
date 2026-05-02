package node

import "k8s.io/api/core/v1"

func NewSlice(
	v []v1.Node,
	cluster string,
) []*Node {
	var result []*Node

	for _, n := range v {
		result = append(result, New(&n, cluster))
	}

	return result
}
