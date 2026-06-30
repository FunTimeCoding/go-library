package trie

func (t *Trie) Insert(token string) {
	node := t.root

	for i := 0; i < len(token); i++ {
		b := token[i]
		child, found := node.children[b]

		if !found {
			child = &Node{children: map[byte]*Node{}}
			node.children[b] = child
		}

		node = child
	}

	node.terminal = true
}
