package trie

func New() *Trie {
	return &Trie{root: &Node{children: map[byte]*Node{}}}
}
