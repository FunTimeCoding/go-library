package rerank

import (
	"fmt"
	"github.com/amikos-tech/pure-onnx/ort"
	"github.com/amikos-tech/pure-tokenizers"
	"github.com/funtimecoding/go-library/pkg/errors"
	"sync"
)

type Reranker struct {
	modelPath      string
	sequenceLength int
	tokenizer      *tokenizers.Tokenizer
	sessions       map[int]*rerankSession
	mu             sync.Mutex
}

func (r *Reranker) Close() error {
	r.mu.Lock()
	defer r.mu.Unlock()

	for _, s := range r.sessions {
		destroySession(s)
	}

	r.sessions = nil

	if r.tokenizer != nil {
		errors.PanicOnError(r.tokenizer.Close())
		r.tokenizer = nil
	}

	return nil
}

func (r *Reranker) Rank(
	query string,
	documents []string,
) ([]Result, error) {
	if len(documents) == 0 {
		return nil, nil
	}

	r.mu.Lock()
	defer r.mu.Unlock()
	batch := len(documents)
	session, e := r.sessionForBatch(batch)

	if e != nil {
		return nil, e
	}

	queries := make([]string, batch)

	for i := range queries {
		queries[i] = query
	}

	encodings, e := r.tokenizer.EncodePairs(
		queries,
		documents,
		tokenizers.WithAddSpecialTokens(),
		tokenizers.WithReturnAttentionMask(),
	)

	if e != nil {
		return nil, fmt.Errorf("tokenize pairs: %w", e)
	}

	for i, enc := range encodings {
		offset := i * r.sequenceLength

		for j := 0; j < r.sequenceLength; j++ {
			if j < len(enc.IDs) {
				session.inputIDs[offset+j] = int64(enc.IDs[j])
			} else {
				session.inputIDs[offset+j] = 0
			}

			if j < len(enc.AttentionMask) {
				session.attentionMask[offset+j] = int64(enc.AttentionMask[j])
			} else {
				session.attentionMask[offset+j] = 0
			}
		}
	}

	if e := session.session.Run(); e != nil {
		return nil, fmt.Errorf("rerank inference: %w", e)
	}

	output := session.outputTensor.GetData()
	results := make([]Result, batch)

	for i := 0; i < batch; i++ {
		results[i] = Result{
			Index: i,
			Score: sigmoid(float64(output[i])),
		}
	}

	return results, nil
}

func (r *Reranker) sessionForBatch(batch int) (*rerankSession, error) {
	if s, okay := r.sessions[batch]; okay {
		return s, nil
	}

	totalTokens := batch * r.sequenceLength
	inputIDs := make([]int64, totalTokens)
	attentionMask := make([]int64, totalTokens)
	inputShape := ort.Shape{int64(batch), int64(r.sequenceLength)}
	inputIDsTensor, e := ort.NewTensor[int64](inputShape, inputIDs)

	if e != nil {
		return nil, fmt.Errorf("create input_ids tensor: %w", e)
	}

	attentionMaskTensor, e := ort.NewTensor[int64](inputShape, attentionMask)

	if e != nil {
		errors.PanicOnError(inputIDsTensor.Destroy())

		return nil, fmt.Errorf("create attention_mask tensor: %w", e)
	}

	outputTensor, e := ort.NewEmptyTensor[float32](ort.Shape{int64(batch), 1})

	if e != nil {
		errors.PanicOnError(attentionMaskTensor.Destroy())
		errors.PanicOnError(inputIDsTensor.Destroy())

		return nil, fmt.Errorf("create output tensor: %w", e)
	}

	session, e := ort.NewAdvancedSession(
		r.modelPath,
		[]string{"input_ids", "attention_mask"},
		[]string{"logits"},
		[]ort.Value{inputIDsTensor, attentionMaskTensor},
		[]ort.Value{outputTensor},
		nil,
	)

	if e != nil {
		errors.PanicOnError(outputTensor.Destroy())
		errors.PanicOnError(attentionMaskTensor.Destroy())
		errors.PanicOnError(inputIDsTensor.Destroy())

		return nil, fmt.Errorf("create rerank session: %w", e)
	}

	s := &rerankSession{
		inputIDs:            inputIDs,
		attentionMask:       attentionMask,
		inputIDsTensor:      inputIDsTensor,
		attentionMaskTensor: attentionMaskTensor,
		outputTensor:        outputTensor,
		session:             session,
	}
	r.sessions[batch] = s

	return s, nil
}
