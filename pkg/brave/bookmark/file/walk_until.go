package file

func WalkUntil(
	n *Node,
	f func(n *Node) bool,
) bool {
	for _, c := range n.Children {
		if f(c) {
			return true
		}

		if WalkUntil(c, f) {
			return true
		}
	}

	return false
}
