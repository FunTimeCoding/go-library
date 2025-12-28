package chroma

import "github.com/amikos-tech/chroma-go/pkg/api/v2"

func (c *Client) CollectionStrings(l v2.Collection) []string {
	var result []string

	for _, d := range c.Get(l).GetDocuments() {
		result = append(result, d.ContentString())
	}

	return result
}
