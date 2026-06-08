package rerank

import (
	"github.com/amikos-tech/pure-tokenizers"
	"sync"
)

type Reranker struct {
	modelPath      string
	sequenceLength int
	tokenizer      *tokenizers.Tokenizer
	sessions       map[int]*rerankSession
	mutex          sync.Mutex
}
