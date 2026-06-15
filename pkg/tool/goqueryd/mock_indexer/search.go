package mock_indexer

import "github.com/funtimecoding/go-library/pkg/face"

func (i *Indexer) Search(
	_ string,
	_ string,
	_ int,
) ([]face.SearchResult, error) {
	return nil, nil
}
