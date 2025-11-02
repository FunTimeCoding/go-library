package node

func SetParents(n *Node) {
	for _, c := range n.Children {
		c.Parent = n
		SetParents(c)
	}
}
