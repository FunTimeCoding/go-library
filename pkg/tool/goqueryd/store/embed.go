package store

import (
	"github.com/funtimecoding/go-library/pkg/generative/ollama"
	"github.com/funtimecoding/go-library/pkg/tool/goqueryd/constant"
	"github.com/funtimecoding/go-library/pkg/tool/goqueryd/store/chunk"
	"github.com/funtimecoding/go-library/pkg/tool/goqueryd/store/embed"
	"time"
)

func (s *Store) Embed(o *ollama.Client) (*embed.Embed, error) {
	pending := s.pendingEmbeddings()
	r := embed.New(len(pending))
	now := time.Now().UTC().Format(time.RFC3339)

	for _, p := range pending {
		chunks := chunk.Document(p.body, p.path)
		texts := make([]string, len(chunks))

		for i, c := range chunks {
			texts[i] = c.Text
		}

		embeddings, e := o.Embed(constant.EmbedModel, texts)

		if e != nil {
			return r, e
		}

		for i, embedding := range embeddings {
			s.insertEmbedding(
				p.hash,
				i,
				chunks[i].Position,
				embedding,
				now,
			)
		}

		r.Chunks += len(chunks)
	}

	return r, nil
}
