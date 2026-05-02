package node

import "k8s.io/api/core/v1"

type Node struct {
	Cluster string
	Name    string
	Raw     *v1.Node
}
