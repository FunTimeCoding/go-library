package node

func FindAllByNameAndType(
	n *Node,
	name string,
	nodeType string,
) []*Node {
	var results []*Node

	Walk(
		n,
		func(o *Node) {
			if o.Name == name && o.Type == nodeType {
				results = append(results, o)
			}
		},
	)

	return results
}
