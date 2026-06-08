package summary_indexer

import (
	"context"
	"github.com/funtimecoding/go-library/pkg/tool/goqueryd/generated/client"
)

func (i *Indexer) Push(
	name string,
	body string,
	metadata map[string]string,
) error {
	merged := map[string]string{"source_type": "session-summary"}

	for key, value := range metadata {
		merged[key] = value
	}

	r, e := i.client.PostDocument(
		context.Background(),
		client.PostDocumentJSONRequestBody{
			Collection: "sessions",
			Path:       name,
			Body:       body,
			Metadata:   &merged,
		},
	)

	if e != nil {
		return e
	}

	return r.Body.Close()
}
