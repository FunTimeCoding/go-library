package store

import (
	"github.com/funtimecoding/go-library/pkg/generative/ollama"
	"github.com/funtimecoding/go-library/pkg/tool/goqueryd/constant"
	"github.com/funtimecoding/go-library/pkg/tool/goqueryd/store/chunk"
	"time"
)

func (s *Store) PushDocument(
	collection string,
	path string,
	body string,
	sourceType string,
	o *ollama.Client,
) error {
	s.ensurePushCollection(collection)
	now := time.Now().UTC().Format(time.RFC3339)
	hash := hashContent(body)
	title := extractTitle(body, path)
	s.insertContent(hash, body, now)
	s.insertDocument(collection, path, title, hash, now)
	chunks := chunk.Document(body, path)
	texts := make([]string, len(chunks))

	for i, c := range chunks {
		texts[i] = c.Text
	}

	embeddings, e := o.Embed(constant.EmbedModel, texts)

	if e != nil {
		return e
	}

	for i, embedding := range embeddings {
		s.insertEmbedding(hash, i, chunks[i].Position, embedding, now)
	}

	if sourceType != "" {
		s.SetSourceType(collection, path, sourceType)
	}

	return nil
}
