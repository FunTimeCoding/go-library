package tokenizer

import "github.com/funtimecoding/go-library/pkg/generative/anthropic/tokenizer/trie"

type Encoder struct {
	trie *trie.Trie
}
