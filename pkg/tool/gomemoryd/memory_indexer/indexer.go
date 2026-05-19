package memory_indexer

import (
	"context"
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/tool/goqueryd/generated/client"
)

type Indexer struct {
	client *client.Client
}

func (i *Indexer) Push(
	path string,
	body string,
) error {
	r, e := i.client.PostDocument(
		context.Background(),
		client.PostDocumentJSONRequestBody{
			Collection: "memories",
			Path:       path,
			Body:       body,
			SourceType: new("memory"),
		},
	)

	if e != nil {
		return e
	}

	return r.Body.Close()
}

func (i *Indexer) MustPush(
	path string,
	body string,
) {
	errors.PanicOnError(i.Push(path, body))
}

func (i *Indexer) Delete(path string) error {
	r, e := i.client.DeleteDocument(
		context.Background(),
		&client.DeleteDocumentParams{Collection: "memories", Path: path},
	)

	if e != nil {
		return e
	}

	return r.Body.Close()
}

func (i *Indexer) MustDelete(path string) {
	errors.PanicOnError(i.Delete(path))
}
