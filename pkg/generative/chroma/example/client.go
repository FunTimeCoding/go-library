package example

import (
	"context"
	"fmt"
	"github.com/amikos-tech/chroma-go/pkg/api/v2"
	"github.com/funtimecoding/go-library/pkg/errors"
)

func Client() {
	c, clientFail := v2.NewHTTPClient()
	errors.PanicOnError(clientFail)
	defer errors.LogClose(c)
	x := context.Background()
	// The default embedding function is used
	l, collectionFail := c.GetOrCreateCollection(
		x,
		"col1",
		v2.WithCollectionMetadataCreate(
			v2.NewMetadata(
				v2.NewStringAttribute("str", "hello"),
				v2.NewIntAttribute("int", 1),
				v2.NewFloatAttribute("float", 1.1),
			),
		),
	)
	errors.PanicOnError(collectionFail)

	errors.PanicOnError(
		l.Add(
			x,
			//v2.WithIDGenerator(v2.NewULIDGenerator()),
			v2.WithIDs("1", "2"),
			v2.WithTexts("hello world", "goodbye world"),
			v2.WithMetadatas(
				v2.NewDocumentMetadata(
					v2.NewIntAttribute("int", 1),
				),
				v2.NewDocumentMetadata(
					v2.NewStringAttribute("str", "hello"),
				),
			),
		),
	)

	n, countFail := l.Count(x)
	errors.PanicOnError(countFail)
	fmt.Printf("Count: %d\n", n)

	r, queryFail := l.Query(x, v2.WithQueryTexts("say hello"))
	errors.PanicOnError(queryFail)

	for _, g := range r.GetDocumentsGroups() {
		for _, d := range g {
			fmt.Printf("Document: %+v\n", d)
		}
	}

	errors.PanicOnError(l.Delete(x, v2.WithIDsDelete("1", "2")))
}
