package node

import core "k8s.io/api/core/v1"

type Node struct {
	Cluster string
	Name    string
	Raw     *core.Node
}
