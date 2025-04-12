package node

import core "k8s.io/api/core/v1"

func New(
	v *core.Node,
	cluster string,
) *Node {
	return &Node{Cluster: cluster, Name: v.Name, Raw: v}
}
