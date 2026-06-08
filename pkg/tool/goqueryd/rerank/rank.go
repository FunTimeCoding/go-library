package rerank

import (
	"fmt"
	"github.com/amikos-tech/pure-tokenizers"
)

func (r *Reranker) Rank(
	query string,
	documents []string,
) ([]Result, error) {
	if len(documents) == 0 {
		return nil, nil
	}

	r.mutex.Lock()
	defer r.mutex.Unlock()
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

	if f := session.session.Run(); f != nil {
		return nil, fmt.Errorf("rerank inference: %w", f)
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
