package file

func Walk(
	n *Node,
	f func(n *Node),
) {
	for _, c := range n.Children {
		f(c)
		Walk(c, f)
	}
}
