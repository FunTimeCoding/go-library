package node

func WalkLevels(
	n *Node,
	maxDepth int,
	f func(n *Node),
) {
	walkLevelsRecursive(n, maxDepth, 0, f)
}
