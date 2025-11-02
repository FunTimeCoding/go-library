package node

func (n *Node) GetPath() []*Node {
	var path []*Node
	current := n

	for current != nil {
		path = append([]*Node{current}, path...)
		current = current.Parent
	}

	return path
}
