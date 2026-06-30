package trie

type Node struct {
	children map[byte]*Node
	terminal bool
}
