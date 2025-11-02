package node

func Collect(n *Node) []*Node {
	var result []*Node

	Walk(
		n,
		func(n *Node) {
			result = append(result, n)
		},
	)

	return result
}
