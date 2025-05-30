package example

import (
	"fmt"
	"github.com/amikos-tech/chroma-go/pkg/api/v2"
	"github.com/funtimecoding/go-library/pkg/generative/chroma"
)

func Client() {
	c := chroma.New()
	defer c.Close()
	// The default embedding function is used
	l := c.Collection(
		"col1",
		v2.WithCollectionMetadataCreate(
			v2.NewMetadata(
				v2.NewStringAttribute("str", "hello"),
				v2.NewIntAttribute("int", 1),
				v2.NewFloatAttribute("float", 1.1),
			),
		),
	)
	c.Add(
		l,
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
	)
	n := c.Count(l)
	fmt.Printf("Count: %d\n", n)

	for _, g := range c.QueryText(l, "say hello").GetDocumentsGroups() {
		for _, d := range g {
			fmt.Printf("Document: %+v\n", d)
		}
	}

	c.DeleteIdentifiers(l, "1", "2")
}
