package summary_indexer

import "github.com/funtimecoding/go-library/pkg/tool/goqueryd/generated/client"

func New(c *client.Client) *Indexer {
	return &Indexer{client: c}
}
