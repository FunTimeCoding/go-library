package file

func walkLevelsRecursive(
	n *Node,
	maxDepth int,
	currentDepth int,
	f func(n *Node),
) {
	if currentDepth >= maxDepth {
		return
	}

	for _, c := range n.Children {
		f(c)
		walkLevelsRecursive(c, maxDepth, currentDepth+1, f)
	}
}
