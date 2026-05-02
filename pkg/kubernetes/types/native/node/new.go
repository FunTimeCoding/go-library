package node

import "k8s.io/api/core/v1"

func New(
	v *v1.Node,
	cluster string,
) *Node {
	return &Node{Cluster: cluster, Name: v.Name, Raw: v}
}
