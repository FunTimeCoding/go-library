package summary_indexer

import (
	"context"
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/tool/goqueryd/generated/client"
)

type Indexer struct {
	client *client.Client
}

func (i *Indexer) Push(
	name string,
	body string,
) error {
	r, e := i.client.PostDocument(
		context.Background(),
		client.PostDocumentJSONRequestBody{
			Collection: "sessions",
			Path:       name,
			Body:       body,
			SourceType: new("session-summary"),
		},
	)

	if e != nil {
		return e
	}

	return r.Body.Close()
}

func (i *Indexer) MustPush(
	name string,
	body string,
) {
	errors.PanicOnError(i.Push(name, body))
}

func (i *Indexer) Delete(path string) error {
	r, e := i.client.DeleteDocument(
		context.Background(),
		&client.DeleteDocumentParams{Collection: "sessions", Path: path},
	)

	if e != nil {
		return e
	}

	return r.Body.Close()
}

func (i *Indexer) MustDelete(path string) {
	errors.PanicOnError(i.Delete(path))
}
